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

package object

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShaSum(t *testing.T) {
	obj, err := FromFile("../../../LICENSE", "blob")
	assert.Nil(t, err)
	sha, err := ShaSum(obj)
	assert.Nil(t, err)
	assert.Equal(t, "d645695673349e3947e8e5ae42332d0ac3164cd7", sha)
}
