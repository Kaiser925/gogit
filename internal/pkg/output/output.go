/*
 * Developed by Kaiser925 on 2021/9/1.
 * Lasted modified 2021/9/1.
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

package output

import (
	"fmt"
	"os"

	"github.com/Kaiser925/gogit/internal/pkg/bytesutil"
)

// Fatal writes data to standard output.
func Fatal(a ...interface{}) {
	fmt.Println(a...)
	os.Exit(1)
}

// HashWriter prints the hash of data.
type HashWriter struct{}

func NewHashWriter() *HashWriter {
	return &HashWriter{}
}

func (h *HashWriter) Write(p []byte) (n int, err error) {
	sha, err := bytesutil.HexSha1(p)
	if err != nil {
		return 0, err
	}
	return os.Stdout.Write([]byte(sha))
}
