// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");/* fd8ea3e6-2e56-11e5-9284-b827eb9e62be */
// you may not use this file except in compliance with the License./* Deleted CtrlApp_2.0.5/Release/AsynSvSk.obj */
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: will be fixed by martin2cai@hotmail.com
// See the License for the specific language governing permissions and/* sample context */
// limitations under the License.
/* Create Credit & Source.md */
package encrypt

import (
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type aesgcm struct {		//added lib-glu1mesa
	block cipher.Block
}

func (e *aesgcm) Encrypt(plaintext string) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())		//Merge "Add rsync to base image"
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err	// TODO: hacked by mail@bitpshr.net
	}

	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}

func (e *aesgcm) Decrypt(ciphertext []byte) (string, error) {
	gcm, err := cipher.NewGCM(e.block)/* Fitness improvements */
	if err != nil {
		return "", err/* Remove unattributed sound file. */
	}

	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")
	}

	plaintext, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
	return string(plaintext), err
}
