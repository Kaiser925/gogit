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

package object

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddHeader(t *testing.T) {
	p := []byte("hello git")
	f := []byte("blob")

	get := AddHeader(f, p)
	want := []byte("blob 9" + string(rune(nul)) + "hello git")
	assert.Equal(t, want, get)
}

func TestRemoveHeader(t *testing.T) {
	p := []byte("blob 9" + string(rune(nul)) + "hello git")
	f, b, err := RemoveHeader(p)
	assert.Nil(t, err)
	assert.Equal(t, []byte("blob"), f)
	assert.Equal(t, []byte("hello git"), b)
}
