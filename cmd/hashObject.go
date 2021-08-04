/*
 * Developed by Kaiser925 on 2021/7/19.
 * Lasted modified 2021/7/19.
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

package cmd

import (
	"fmt"
	"os"

	"github.com/Kaiser925/gogit/internal/pkg/repository"
	"github.com/Kaiser925/gogit/internal/pkg/repository/object"
	"github.com/spf13/cobra"
)

// hashObjectCmd represents the catFile command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-object",
	Short: "Compute object ID and optionally creates a blob from a file",
	Run:   run,
}

var (
	objType   string
	needWrite bool
)

func init() {
	rootCmd.AddCommand(hashObjectCmd)
	flags := hashObjectCmd.Flags()
	flags.StringVarP(&objType, "type", "t", "blob", "Specify the type (default: \"blob\")")
	flags.BoolVar(&needWrite, "w", false, "Actually write the object into the object database")
}

func run(_ *cobra.Command, args []string) {
	if !object.IsValid(objType) {
		println("not valid type:", objType)
		os.Exit(1)
	}

	if len(args) == 0 {
		return
	}

	obj, err := object.FromFile(args[0], objType)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if !needWrite {
		sha, err := object.ShaSum(obj)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		fmt.Println(sha)
		return
	}

	repo, err := repository.New(".")
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	err = repo.WriteObject(obj)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
