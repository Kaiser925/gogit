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

type ObjFunc func([]byte) Object

var _objMap = map[string]ObjFunc{
	"blob":   NewBlob,
	"commit": NewCommit,
	"tree":   NewTree,
	"tag":    NewTag,
}

// Unmarshal reads from r, parses data to Object.
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

// ReadFile reads Object from file.
func ReadFile(name string) (Object, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Unmarshal(f)
}

// MarshalTo encodes Object to data, writes it to w.
func MarshalTo(obj Object, w io.Writer) error {
	p, err := obj.Serialize()
	if err != nil {
		return err
	}
	size := strconv.Itoa(len(p))
	h := sha1.New()
	h.Write(obj.Format())
	h.Write([]byte{32})
	h.Write([]byte(size))
	h.Write([]byte{0x00})
	h.Write(p)
	res := h.Sum(nil)

	hw := hex.NewEncoder(w)
	_, err = hw.Write(res)
	if err != nil {
		return err
	}
	return nil
}

// WriteTo decodes Object and write it to file,
func WriteTo(obj Object, name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	defer f.Close()
	return MarshalTo(obj, f)
}
