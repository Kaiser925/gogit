/*
 * Developed by Kaiser925 on 2021/8/31.
 * Lasted modified 2021/8/31.
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

package bytesutil

import (
	"fmt"
)

func ExampleHexSha1() {
	src := []byte("Hello")
	encodedStr, _ := HexSha1(src)
	fmt.Printf("%s\n", encodedStr)
	// output:
	// f7ff9e8b7bb2e09b70935a5d785e0cc5d9d0abf0
}
