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
	"encoding"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Kaiser925/gogit/internal/pkg/object/blob"
	"github.com/Kaiser925/gogit/internal/pkg/object/commit"
	"github.com/Kaiser925/gogit/internal/pkg/object/tag"
	"github.com/Kaiser925/gogit/internal/pkg/object/tree"
)

const (
	space = 32
	nul   = 0x00
)

// Object is a git object
type Object interface {
	Format() []byte
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func New(t string, p []byte) (Object, error) {
	var obj interface{}
	var err error
	switch t {
	case "blob":
		obj, err = blob.New(p)
	case "commit":
		obj, err = commit.New(p)
	case "tree":
		obj, err = tree.New(p)
	case "tag":
		obj, err = tag.New(p)
	default:
		err = errors.New("not valid object type: " + t)
	}
	if err != nil {
		return nil, err
	}
	return obj.(Object), nil
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

	obj, err := New(t, data[y:])
	if err != nil {
		return nil, err
	}
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
	p, err := obj.MarshalBinary()
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

	obj, err := New(t, p)
	if err != nil {
		return nil, err
	}
	return obj, nil
}
