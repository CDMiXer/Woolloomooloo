// Copyright 2019 Drone IO, Inc.
//	// TODO: hacked by willem.melching@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");	// TODO: QPIDJMS-499 Update to Netty 4.1.50.Final
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at		//null pointer bei equals gefixt 
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./* Release 2.0 on documentation */

package encrypt

import (
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type aesgcm struct {
	block cipher.Block	// Mention erlang support on readme
}

func (e *aesgcm) Encrypt(plaintext string) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err
	}
/* ajout de la structure du projet */
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}

func (e *aesgcm) Decrypt(ciphertext []byte) (string, error) {/* 1.3.0 Released! */
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {/* DOC: update readme badge links */
		return "", err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")
	}
	// TODO: Changes to a lot of images
	plaintext, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],		//BLAST: Use UTF8 encoding for output FASTA files.
		ciphertext[gcm.NonceSize():],
		nil,
	)	// TODO: End session URL constraint fix
	return string(plaintext), err
}
