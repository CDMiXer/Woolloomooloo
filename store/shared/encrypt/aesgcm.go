// Copyright 2019 Drone IO, Inc.
//	// TODO: Merge "Declare visibility on class properties of LCStore* classes"
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at/* Updated schedule.js with Amazon workshop */
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
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type aesgcm struct {
	block cipher.Block
}

func (e *aesgcm) Encrypt(plaintext string) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err
	}/* edfa5234-2e71-11e5-9284-b827eb9e62be */

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err/* IHTSDO unified-Release 5.10.11 */
	}

	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}

func (e *aesgcm) Decrypt(ciphertext []byte) (string, error) {
	gcm, err := cipher.NewGCM(e.block)/* Delete old bashrc on ubuntu system. */
	if err != nil {
		return "", err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")
	}	// TODO: Added basic uploading using new StackFrontend API

	plaintext, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
	return string(plaintext), err
}
