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
	"encoding"
	"errors"
	"io"
	"os"
	"strconv"
)

const (
	space = 32
	nul   = 0x00
)

// GitObject is the interface that wraps the basic object method.
type GitObject interface {
	Format() []byte

	// String returns string of GitObject content.
	String() string
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

func New(t string, p []byte) (GitObject, error) {
	var obj interface{}
	var err error
	switch t {
	case "blob":
		obj, err = NewBlob(p)
	case "commit":
		obj, err = NewCommit(p)
	case "tree":
		obj, err = NewTree(p)
	case "tag":
		obj, err = NewTag(p)
	default:
		err = errors.New("not valid object type: " + t)
	}
	if err != nil {
		return nil, err
	}
	return obj.(GitObject), nil
}

// removeHeader remove the header from data, return the format, data and error.
func removeHeader(p []byte) ([]byte, []byte, error) {
	x := bytes.Index(p, []byte{space})
	f := p[0:x]

	data := p[x+1:]
	// get format
	y := bytes.Index(data, []byte{nul})

	size, err := strconv.Atoi(string(data[0:y]))
	if err != nil {
		return nil, nil, err
	}

	if size != len(data)-y-1 {
		return nil, nil, errors.New("bad length for sha")
	}

	return f, data[y+1:], nil
}

// addHeader inserts the header to object data.
func addHeader(format []byte, p []byte) []byte {
	buf := bytes.NewBuffer([]byte{})
	size := strconv.Itoa(len(p))
	buf.Write(format)
	buf.Write([]byte{space})
	buf.Write([]byte(size))
	buf.Write([]byte{nul})
	buf.Write(p)
	return buf.Bytes()
}

// Convert converts the file to GitObject.
func Convert(name string, t string) (GitObject, error) {
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

// Parse parses the binary data to specify GitObject.
func Parse(t string, p []byte) (GitObject, error) {
	var o GitObject
	switch t {
	case "blob":
		o = &Blob{}
	default:
		return nil, errors.New("not valid object type: " + t)
	}

	err := o.UnmarshalBinary(p)
	if err != nil {
		return nil, err
	}
	return o, nil
}
