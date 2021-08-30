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

package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/Kaiser925/gogit/internal/pkg/repository"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init an empty Git repository or reinitialize an existing one",
	Run: func(cmd *cobra.Command, args []string) {
		dir := "."
		if len(args) > 0 {
			dir = args[0]
		}
		_, err := repository.Init(dir)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		name, _ := filepath.Abs(dir)
		fmt.Println("Initialized Git repository in", path.Join(name, ".git/"))
	},
}
