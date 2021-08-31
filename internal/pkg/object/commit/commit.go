/*
 * Developed by Kaiser925 on 2021/8/31.
 * Lasted modified 2021/8/4.
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

package commit

type Commit struct {
	p []byte
}

func New(p []byte) (*Commit, error) {
	c := &Commit{}
	err := c.UnmarshalBinary(p)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Commit) Format() []byte {
	return []byte("commit")
}

func (c *Commit) MarshalBinary() ([]byte, error) {
	panic("implement me")
}

func (c *Commit) UnmarshalBinary(bytes []byte) error {
	panic("implement me")
}
