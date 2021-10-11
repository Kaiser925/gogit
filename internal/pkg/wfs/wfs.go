/*
 * Developed by Kaiser925 on 2021/10/11.
 * Lasted modified 2021/10/11.
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

package wfs

import (
	"io/fs"
	"os"
)

// A WriteFS is a file system that can write file.
type WriteFS interface {
	OpenFile(name string, flag int, perm fs.FileMode) (WFile, error)
}

type WFile interface {
	Stat() (fs.FileInfo, error)
	Write(p []byte) (n int, err error)
	Close() error
}

type DirWriteFS interface {
	WriteFS
	Mkdir(name string, perm fs.FileMode) error
	MkdirAll(name string, perm fs.FileMode) error
}

// Create creates or truncates the named file. If the file already exists,
// it is truncated. If the file does not exist, it is created with mode 0666
// (before umask).
func Create(fsys WriteFS, name string) (WFile, error) {
	return fsys.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

// WriteFile writes data to WFS file.
func WriteFile(fsys WriteFS, name string, data []byte) (int, error) {
	f, err := fsys.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return 0, err
	}
	return f.Write(data)
}

type wFS string

// WFS returns a file system (an WriteFS) for the tree of files rooted at the directory dir.
func WFS(name string) DirWriteFS {
	return wFS(name)
}

func (f wFS) OpenFile(name string, flag int, perm fs.FileMode) (WFile, error) {
	file, err := os.OpenFile(string(f)+"/"+name, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (f wFS) Mkdir(name string, perm fs.FileMode) error {
	return os.Mkdir(string(f)+"/"+name, perm)
}

func (f wFS) MkdirAll(name string, perm fs.FileMode) error {
	return os.MkdirAll(string(f)+"/"+name, perm)
}
