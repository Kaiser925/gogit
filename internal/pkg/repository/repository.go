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
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/Kaiser925/gogit/internal/pkg/object"

	"github.com/Kaiser925/gogit/internal/pkg/file"

	"gopkg.in/ini.v1"
)

type repository struct {
	workTree string
	gitDir   string

	conf *ini.File
}

// New returns a repository instance.
func New(p string) (*repository, error) {
	var r repository
	r.workTree = p
	r.gitDir = path.Join(p, ".git")

	f, err := os.Stat(r.gitDir)
	if err != nil {
		return nil, fmt.Errorf("construct repository error: %w", err)
	}
	if !f.IsDir() {
		return nil, errors.New(fmt.Sprintf("%s is not a git repository", r.gitDir))
	}

	conf, err := ini.Load(path.Join(r.gitDir, "config"))
	if err != nil {
		return nil, err
	}
	r.conf = conf
	return &r, nil
}

// WriteObject writes object to repository object database
func (r *repository) WriteObject(obj object.Object) error {
	data, err := object.Marshal(obj)
	if err != nil {
		return err
	}

	sha, err := object.ShaSum(obj)
	if err != nil {
		return err
	}
	if err := file.MkdirOr(r.gitDir, sha[:2]); err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(r.gitDir, sha[:2], sha[2:]))
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

// Init inits a new git repository.
func Init(p string) (*repository, error) {
	var r repository
	r.workTree = p
	r.gitDir = path.Join(p, ".git")
	dirs := []string{"branches", "objects", "refs/tags", "refs/heads"}
	for _, d := range dirs {
		err := file.MkdirOr(r.gitDir, d)
		if err != nil {
			return nil, err
		}
	}

	desc := []byte("Unnamed repository; edit this file 'description' to name the repository.\n")
	err := file.WriteTo(path.Join(r.gitDir, "description"), desc)
	if err != nil {
		return nil, err
	}

	ref := []byte("ref: refs/heads/master\n")
	err = file.WriteTo(path.Join(r.gitDir, "HEAD"), ref)
	if err != nil {
		return nil, err
	}

	confName := path.Join(r.gitDir, "config")
	err = file.WriteTo(confName, defaultConf())
	if err != nil {
		return nil, err
	}

	conf, err := ini.Load(confName)
	if err != nil {
		return nil, err
	}
	r.conf = conf
	if err != nil {
		return nil, err
	}
	return &r, nil
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
