/*
 * Developed by Kaiser925 on 2021/9/28.
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

package object

type Tree struct {
	p []byte
}

func NewTree(p []byte) (*Tree, error) {
	t := &Tree{}
	err := t.UnmarshalBinary(p)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Tree) Format() []byte {
	return []byte("tree")
}

func (c *Tree) MarshalBinary() ([]byte, error) {
	panic("implement me")
}

func (c *Tree) UnmarshalBinary(bytes []byte) error {
	panic("implement me")
}

func (c *Tree) String() string {
	panic("implement me")
}
