/*
 * Developed by Kaiser925 on 2021/7/13.
 * Lasted modified 2021/7/13.
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

package repository

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/Kaiser925/gogit/internal/pkg/wfs"

	"github.com/Kaiser925/gogit/internal/pkg/bytesutil"

	"gopkg.in/ini.v1"
)

var ErrNotDir = errors.New("not a dir")

type repository struct {
	conf *ini.File

	fsys  fs.FS
	fsysw wfs.DirWriteFS
}

func NewWithFS(fsys fs.FS, fsysw wfs.DirWriteFS) (*repository, error) {
	f, err := fsys.Open(".")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, ErrNotDir
	}
	b, err := fs.ReadFile(fsys, "config")
	if err != nil {
		return nil, err
	}
	conf, err := ini.Load(b)
	if err != nil {
		return nil, err
	}
	return &repository{
		fsys:  fsys,
		fsysw: fsysw,
		conf:  conf,
	}, nil
}

// New returns a repository instance.
func New(p string) (*repository, error) {
	f := os.DirFS(p + "/.git")
	wf := wfs.WFS(p + "/.git")
	return NewWithFS(f, wf)
}

// Write implements io.Writer.
// Write writes the GitObject binary data to file system.
func (r *repository) Write(p []byte) (n int, err error) {
	sha, err := bytesutil.HexSha1(p)
	if err != nil {
		return 0, err
	}

	name := "objects/" + sha[:2]
	if _, err := fs.Stat(r.fsys, name); err != nil {
		err = r.fsysw.Mkdir(name, os.ModePerm)
		if err != nil {
			return 0, err
		}
	}

	f, err := wfs.Create(r.fsysw, filepath.Join("objects", sha[:2], sha[2:]))
	if err != nil {
		return 0, err
	}
	defer f.Close()

	d, err := bytesutil.ZlibCompress(p)
	if err != nil {
		return
	}
	n, err = f.Write(d)
	if err != nil {
		return
	}
	return
}

// Open opens the sha file from git file system.
func (r *repository) Open(sha string) ([]byte, error) {
	p, err := fs.ReadFile(r.fsys, path.Join("objects", sha[0:2], sha[2:]))
	if err != nil {
		return nil, err
	}
	p, err = bytesutil.ZlibDecompress(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Init inits a new git repository.
func Init(p string) (*repository, error) {
	return InitWithFs(os.DirFS(p+"/.git"), wfs.WFS(p+"/.git"))
}

func InitWithFs(fsys fs.FS, fsysw wfs.DirWriteFS) (*repository, error) {
	r := &repository{
		fsys:  fsys,
		fsysw: fsysw,
	}

	dirs := []string{"branches", "objects", "refs/tags", "refs/heads"}
	for _, d := range dirs {
		if _, err := fs.Stat(r.fsys, d); err != nil {
			if err := r.fsysw.Mkdir(d, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}

	desc := []byte("Unnamed repository; edit this file 'description' to name the repository.\n")
	_, err := wfs.WriteFile(r.fsysw, "description", desc)
	if err != nil {
		return nil, err
	}

	ref := []byte("ref: refs/heads/master\n")
	_, err = wfs.WriteFile(r.fsysw, "HEAD", ref)
	if err != nil {
		return nil, err
	}

	_, err = wfs.WriteFile(r.fsysw, "config", ref)
	if err != nil {
		return nil, err
	}

	conf, err := ini.Load(defaultConf())
	if err != nil {
		return nil, err
	}
	r.conf = conf
	if err != nil {
		return nil, err
	}
	return r, nil
}

func defaultConf() []byte {
	return []byte(`
[core]
        repositoryformatversion = 0
        filemode = true
        bare = false
        logallrefupdates = true
        ignorecase = true
        precomposeunicode = true`)
}
