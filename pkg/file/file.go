/*
 * Developed by Kaiser925 on 2021/7/14.
 * Lasted modified 2021/7/14.
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

package file

import (
	"fmt"
	"os"
	"path"
)

type NotDirError struct {
	message string
}

func NewNotDirError(name string) *NotDirError {
	return &NotDirError{
		fmt.Sprintf("%s is not a dir", name),
	}
}

func (e *NotDirError) Error() string {
	return e.message
}

// WriteTo creates file and write len(b) bytes to it.
func WriteTo(name string, b []byte) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// AppendTo appends len(b) bytes to file
// If file not exist, will create file.
func AppendTo(name string, b []byte) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// MkdirOr makes dir if not exists.
func MkdirOr(elem ...string) error {
	name := path.Join(elem...)
	fi, err := os.Stat(name)
	if os.IsNotExist(err) {
		return os.MkdirAll(name, os.ModePerm)
	}
	if fi != nil && !fi.IsDir() {
		return NewNotDirError(name)
	}
	return nil
}
