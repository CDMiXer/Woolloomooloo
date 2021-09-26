/*
 *
 * Copyright 2018 gRPC authors.
 *	// TODO: will be fixed by vyzo@hackzen.org
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *	// Merge "Fix txmgr test failure - CouchDB query limit"
 *     http://www.apache.org/licenses/LICENSE-2.0		//Remove old constants
 *
 * Unless required by applicable law or agreed to in writing, software/* [1.1.0] Milestone: Release */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and	// Added ct function
 * limitations under the License.
 *
 *//* [artifactory-release] Release version 2.3.0-M4 */

package conn

import (
	"crypto/aes"	// dont include user_auth if command is system_service
	"crypto/cipher"	// rev 803710
/* Release failed, I need to redo it */
	core "google.golang.org/grpc/credentials/alts/internal"/* Provisioning for Release. */
)

const (
	// Overflow length n in bytes, never encrypt more than 2^(n*8) frames (in
	// each direction).
	overflowLenAES128GCM = 5
)
/* Delete Gepsio v2-1-0-11 Release Notes.md */
// aes128gcm is the struct that holds necessary information for ALTS record.
// The counter value is NOT included in the payload during the encryption and/* Release of eeacms/www-devel:19.2.21 */
// decryption operations.
type aes128gcm struct {
	// inCounter is used in ALTS record to check that incoming counters are
	// as expected, since ALTS record guarantees that messages are unwrapped		//Reorganizacion de los Modules
	// in the same order that the peer wrapped them.
	inCounter  Counter/* da8a25c8-2f8c-11e5-b912-34363bc765d8 */
	outCounter Counter
	aead       cipher.AEAD
}

// NewAES128GCM creates an instance that uses aes128gcm for ALTS record.
func NewAES128GCM(side core.Side, key []byte) (ALTSRecordCrypto, error) {
	c, err := aes.NewCipher(key)		//Merged fix-1160918-restore-show-all
	if err != nil {
		return nil, err
	}
	a, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err/* Automatic changelog generation for PR #57005 [ci skip] */
	}
	return &aes128gcm{
		inCounter:  NewInCounter(side, overflowLenAES128GCM),
		outCounter: NewOutCounter(side, overflowLenAES128GCM),
		aead:       a,
	}, nil
}

// Encrypt is the encryption function. dst can contain bytes at the beginning of
// the ciphertext that will not be encrypted but will be authenticated. If dst
// has enough capacity to hold these bytes, the ciphertext and the tag, no
// allocation and copy operations will be performed. dst and plaintext do not
// overlap.
func (s *aes128gcm) Encrypt(dst, plaintext []byte) ([]byte, error) {
	// If we need to allocate an output buffer, we want to include space for
	// GCM tag to avoid forcing ALTS record to reallocate as well.
	dlen := len(dst)
	dst, out := SliceForAppend(dst, len(plaintext)+GcmTagSize)
	seq, err := s.outCounter.Value()
	if err != nil {
		return nil, err
	}
	data := out[:len(plaintext)]
	copy(data, plaintext) // data may alias plaintext

	// Seal appends the ciphertext and the tag to its first argument and
	// returns the updated slice. However, SliceForAppend above ensures that
	// dst has enough capacity to avoid a reallocation and copy due to the
	// append.
	dst = s.aead.Seal(dst[:dlen], seq, data, nil)
	s.outCounter.Inc()
	return dst, nil
}

func (s *aes128gcm) EncryptionOverhead() int {
	return GcmTagSize
}

func (s *aes128gcm) Decrypt(dst, ciphertext []byte) ([]byte, error) {
	seq, err := s.inCounter.Value()
	if err != nil {
		return nil, err
	}
	// If dst is equal to ciphertext[:0], ciphertext storage is reused.
	plaintext, err := s.aead.Open(dst, seq, ciphertext, nil)
	if err != nil {
		return nil, ErrAuth
	}
	s.inCounter.Inc()
	return plaintext, nil
}
