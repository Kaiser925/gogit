/*
 * Developed by Kaiser925 on 2021/8/31.
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

package bytesutil

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

// HexSha1 calculates the SHA1 sum of data,
// and returns the hexadecimal encoding of sum
func HexSha1(p []byte) (string, error) {
	sh := sha1.New()
	_, err := sh.Write(p)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sh.Sum(nil)), nil
}

// ZlibCompress compresses p with zlib.
func ZlibCompress(p []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	_, err := w.Write(p)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ZlibDecompress decompresses the data.
func ZlibDecompress(p []byte) ([]byte, error) {
	buf := bytes.NewReader(p)
	r, err := zlib.NewReader(buf)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}
