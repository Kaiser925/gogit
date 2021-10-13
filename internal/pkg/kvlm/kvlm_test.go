/*
 * Developed by Kaiser925 on 2021/10/13.
 * Lasted modified 2021/10/13.
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

package kvlm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var parseTests = []struct {
	raw  []byte
	want map[string][]string
}{
	{
		raw: []byte(`tree tree

Body
`),
		want: map[string][]string{
			"tree": []string{"tree"},
			Body:   []string{"Body\n"},
		},
	},
	{
		raw: []byte(`tree tree

Body`),
		want: map[string][]string{
			"tree": []string{"tree"},
			Body:   []string{"Body"},
		},
	},
	{
		raw: []byte(`tree tree`),
		want: map[string][]string{
			"tree": []string{"tree"},
		},
	},
	{
		raw: []byte(`header1 header1
header2 header2line1

 header2line2
 header2line3

Bodyline1
Bodyline2

Bodyline3
`),
		want: map[string][]string{
			"header1": []string{"header1"},
			"header2": []string{"header2line1\n\n header2line2\n header2line3\n"},
			Body:      []string{"Bodyline1\nBodyline2\n\nBodyline3\n"},
		},
	},
}

func TestParse(t *testing.T) {
	for _, tt := range parseTests {
		assert.Equal(t, tt.want, Parse(tt.raw).kvs)
	}
}
