/*
 * Developed by Kaiser925 on 2021/9/29.
 * Lasted modified 2021/9/29.
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

// Package kvlm is used to handle "Key-Value List with Message"
package kvlm

import (
	"bufio"
	"bytes"
	"strings"
)

type KVListMap map[string][]string

const Body = " "

// Parse parses the raw data to KVListMap
func Parse(raw []byte) KVListMap {
	listMap := make(KVListMap)

	var endString string
	last := len(raw) - 1
	if raw[last] == '\n' {
		endString = "\n"
	}

	// split raw to lines.
	lines := make([]string, 0, 32)
	scan := bufio.NewScanner(bytes.NewReader(raw))
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}

	var prevKey string
	bodyIndex := -1
	for i, line := range lines {
		// If current line is empty line, peek next line.
		// If next line is start with " ", add "\n" to previous key.
		// else
		if line == "" {
			// no more lines.
			if i+1 >= len(lines) {
				break
			}

			if strings.Index(lines[i+1], " ") == 0 {
				listMap[prevKey][0] = listMap[prevKey][0] + "\n\n"
				continue
			}

			// rest of lines is body.
			bodyIndex = i + 1
			break
		}

		space := strings.Index(line, " ")
		// line start with " " means it is part of multiple lines value.
		if space == 0 {
			strs := listMap[prevKey]
			last := len(strs) - 1
			strs[last] = strs[last] + line + "\n"
			continue
		}

		// load headers.
		if space > 0 {
			key := line[:space]
			_, ok := listMap[key]
			if ok {
				listMap[key] = append(listMap[key], line)
			} else {
				listMap[key] = []string{line[space+1:]}
			}
			prevKey = key
		}
	}

	if bodyIndex > 0 {
		listMap[Body] = []string{strings.Join(lines[bodyIndex:], "\n")}
		listMap[Body][0] = listMap[Body][0] + endString
	}

	return listMap
}
