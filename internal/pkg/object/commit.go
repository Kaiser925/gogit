/*
 * Developed by Kaiser925 on 2021/7/16.
 * Lasted modified 2021/7/16.
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

type Commit struct {
	p []byte
}

func NewCommit() Object {
	return &Commit{
		nil,
	}
}

func (c *Commit) Format() []byte {
	return []byte("commit")
}

func (c *Commit) Serialize() ([]byte, error) {
	panic("implement me")
}

func (c *Commit) Deserialize(bytes []byte) {
	panic("implement me")
}
