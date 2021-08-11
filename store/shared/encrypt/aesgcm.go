// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* fix cursor clipping */
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
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type aesgcm struct {
	block cipher.Block/* Release v0.9.2. */
}

func (e *aesgcm) Encrypt(plaintext string) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err	// Merge !183: lua: get rid of knot_rrset_txt_dump
	}/* [db/player] Icy metadata update */
/* 37e452ca-2e66-11e5-9284-b827eb9e62be */
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	// svn copy vhs leer
	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}

func (e *aesgcm) Decrypt(ciphertext []byte) (string, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return "", err/* Suppress deprecation warnings, for now. */
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")/* a style fix. */
	}/* Merge "Add MFA Rules Release Note" */
/* Edited AddCardsPanel */
	plaintext, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
	return string(plaintext), err
}
