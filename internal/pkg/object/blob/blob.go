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

package blob

// Blob represents a git blob
type Blob struct {
	p []byte
}

func New(p []byte) (*Blob, error) {
	blob := &Blob{}
	err := blob.UnmarshalBinary(p)
	if err != nil {
		return nil, err
	}
	return blob, nil
}

func (b *Blob) Format() []byte {
	return []byte("blob")
}

func (b *Blob) MarshalBinary() ([]byte, error) {
	return b.p, nil
}

func (b *Blob) UnmarshalBinary(p []byte) error {
	b.p = p
	return nil
}
