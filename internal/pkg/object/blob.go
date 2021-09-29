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

// Blob represents a git blob
type Blob struct {
	p []byte
}

func NewBlob(p []byte) (*Blob, error) {
	return &Blob{p: p}, nil
}

func (b *Blob) Format() []byte {
	return []byte("blob")
}

func (b *Blob) MarshalBinary() ([]byte, error) {
	return AddHeader(b.Format(), b.p), nil
}

func (b *Blob) UnmarshalBinary(p []byte) error {
	_, p, err := RemoveHeader(p)
	if err != nil {
		return err
	}
	b.p = p
	return nil
}
