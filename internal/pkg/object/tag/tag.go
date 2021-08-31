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

package tag

import "github.com/Kaiser925/gogit/internal/pkg/object"

type Tag struct {
	p []byte
}

func New(p []byte) object.Object {
	c := &Tag{}
	c.Deserialize(p)
	return c
}

func (c *Tag) Format() []byte {
	return []byte("tag")
}

func (c *Tag) Serialize() ([]byte, error) {
	panic("implement me")
}

func (c *Tag) Deserialize(bytes []byte) {
	panic("implement me")
}
