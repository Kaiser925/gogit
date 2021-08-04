/*
 * Developed by Kaiser925 on 2021/7/16.
 * Lasted modified 2021/7/16.
 * Copyright (c) 2021.  All rights reserved
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package object

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	space = 32
	nul   = 0x00
)

// Object is a git object
type Object interface {
	Format() []byte
	Serialize() ([]byte, error)
	Deserialize([]byte)
}

type NewObjFunc func([]byte) Object

var _objMap = map[string]NewObjFunc{
	"blob":   NewBlob,
	"commit": NewCommit,
	"tree":   NewTree,
	"tag":    NewTag,
}

// IsValid checks t is valid object type or not.
func IsValid(t string) bool {
	_, ok := _objMap[t]
	return ok
}

// Unmarshal reads data from reader of object database, parses data to Object.
func Unmarshal(r io.Reader) (Object, error) {
	rc, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	p, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	x := bytes.Index(p, []byte{space})
	t := string(p[0:x])

	data := p[x+1:]
	y := bytes.Index(data, []byte{nul})

	size, err := strconv.Atoi(string(data[0:y]))
	if err != nil {
		return nil, err
	}

	if size != len(data)-y-1 {
		return nil, errors.New("bad length for sha")
	}

	objFunc, ok := _objMap[t]
	if !ok {
		return nil, errors.New(fmt.Sprintf("unknown type %s", t))
	}
	obj := objFunc(data[y:])
	return obj, nil
}

// Load loads Object from object file.
func Load(name string) (Object, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Unmarshal(f)
}

// Marshal encodes Object to bytes.
func Marshal(obj Object) ([]byte, error) {
	p, err := obj.Serialize()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer([]byte{})

	size := strconv.Itoa(len(p))
	buf.Write(obj.Format())
	buf.Write([]byte{space})
	buf.Write([]byte(size))
	buf.Write([]byte{nul})
	buf.Write(p)
	return buf.Bytes(), nil
}

// FromFile computes file, returns the object.
func FromFile(name string, t string) (Object, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	newObj, ok := _objMap[t]
	if !ok {
		return nil, errors.New(fmt.Sprintf("unknown type %s", t))
	}

	obj := newObj(p)
	return obj, nil
}

// ShaSum generates the SHA value of object, encodes sum to hex string.
func ShaSum(obj Object) (string, error) {
	hasher := sha1.New()
	data, err := Marshal(obj)
	if err != nil {
		return "", err
	}
	_, err = hasher.Write(data)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
