/*
 * Developed by Kaiser925 on 2021/9/28.
 * Lasted modified 2021/9/28.
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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHexSha1(t *testing.T) {
	tests := []struct {
		val  []byte
		want string
	}{
		{
			val:  []byte("hello"),
			want: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
		{
			val:  []byte("ok"),
			want: "7a85f4764bbd6daf1c3545efbbf0f279a6dc0beb",
		},
	}

	for _, tt := range tests {
		got, err := HexSha1(tt.val)
		assert.Nil(t, err)
		assert.Equal(t, tt.want, got)
	}
}

func TestZlibCompressAndDecompress(t *testing.T) {
	txt := []byte("hello, zlib")
	p, err := ZlibCompress(txt)
	assert.Nil(t, err)

	b, err := ZlibDecompress(p)
	assert.Nil(t, err)
	assert.Equal(t, txt, b)
}
