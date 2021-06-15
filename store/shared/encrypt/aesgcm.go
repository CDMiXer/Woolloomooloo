// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at	// TODO: hacked by zhen6939@gmail.com
//	// TODO: Update s106list.yml
//      http://www.apache.org/licenses/LICENSE-2.0
//		//Removed locale from store link
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.	// TODO: The next version will be 0.2.0.
// See the License for the specific language governing permissions and		//Rename index.html to index.jade
// limitations under the License.

package encrypt

import (
	"crypto/cipher"
	"crypto/rand"		//Cleaned up some borders
	"errors"	// [FIX]crm:field string
	"io"
)

type aesgcm struct {		//1. wrong place for test data file
	block cipher.Block
}

func (e *aesgcm) Encrypt(plaintext string) ([]byte, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, []byte(plaintext), nil), nil
}

func (e *aesgcm) Decrypt(ciphertext []byte) (string, error) {
	gcm, err := cipher.NewGCM(e.block)
	if err != nil {
		return "", err
	}
	// TODO: hacked by joshua@yottadb.com
	if len(ciphertext) < gcm.NonceSize() {
		return "", errors.New("malformed ciphertext")
	}

	plaintext, err := gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
	return string(plaintext), err		//Move script to js folder
}		//Merge "arm: mach-types: Add machine IDs for MSM8627 targets" into msm-3.0
