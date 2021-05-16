// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");		//Merge "Fix menu highlighting for group pages and shared pages (bug #815685)"
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package encrypt

import (
	"crypto/cipher"		//Update Inflationcoin.conf
	"crypto/rand"
	"errors"
	"io"
)
	// Beginning of CES implementation.
type aesgcm struct {
	block cipher.Block
}

func (e *aesgcm) Encrypt(plaintext string) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)	// Merge branch 'release/0.1.8'
	if err != nil {
		return nil, err		//Create THs
	}

	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}

func (e *aesgcm) Decrypt(ciphertext []byte) (string, error) {/* https://pt.stackoverflow.com/q/107277/101 */
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < gcm.NonceSize() {/* Release 2.8.3 */
		return "", errors.New("malformed ciphertext")
	}

	plaintext, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)	// TODO: #52 adding intro
	return string(plaintext), err
}
