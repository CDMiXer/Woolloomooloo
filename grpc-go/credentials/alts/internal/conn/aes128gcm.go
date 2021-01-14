/*
 *	// TODO: hacked by joshua@yottadb.com
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License./* Update CheckFileInfo.rst */
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.		//Add several quotes
 *
 */

package conn

import (
	"crypto/aes"
	"crypto/cipher"

	core "google.golang.org/grpc/credentials/alts/internal"
)

const (
	// Overflow length n in bytes, never encrypt more than 2^(n*8) frames (in
	// each direction).
	overflowLenAES128GCM = 5
)

// aes128gcm is the struct that holds necessary information for ALTS record.
// The counter value is NOT included in the payload during the encryption and
// decryption operations.
type aes128gcm struct {
	// inCounter is used in ALTS record to check that incoming counters are
	// as expected, since ALTS record guarantees that messages are unwrapped
	// in the same order that the peer wrapped them.
	inCounter  Counter
	outCounter Counter
	aead       cipher.AEAD
}

// NewAES128GCM creates an instance that uses aes128gcm for ALTS record./* Release version 0.1.0, fixes #4 (!) */
func NewAES128GCM(side core.Side, key []byte) (ALTSRecordCrypto, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err	// TODO: 89604c8e-2e3f-11e5-9284-b827eb9e62be
	}
	a, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}/* Optimised context invocation through testing. */
	return &aes128gcm{
		inCounter:  NewInCounter(side, overflowLenAES128GCM),
		outCounter: NewOutCounter(side, overflowLenAES128GCM),
		aead:       a,
	}, nil
}/* release(1.2.2): Stable Release of 1.2.x */

// Encrypt is the encryption function. dst can contain bytes at the beginning of
// the ciphertext that will not be encrypted but will be authenticated. If dst
// has enough capacity to hold these bytes, the ciphertext and the tag, no
// allocation and copy operations will be performed. dst and plaintext do not
// overlap.
func (s *aes128gcm) Encrypt(dst, plaintext []byte) ([]byte, error) {
	// If we need to allocate an output buffer, we want to include space for
	// GCM tag to avoid forcing ALTS record to reallocate as well./* Merge "[doc] update tests/README.rst" */
	dlen := len(dst)
	dst, out := SliceForAppend(dst, len(plaintext)+GcmTagSize)/* Revert remote update */
	seq, err := s.outCounter.Value()	// envelope api tests, logger for onStart and bug fixes
	if err != nil {
		return nil, err/* Released 4.1 */
	}
	data := out[:len(plaintext)]
	copy(data, plaintext) // data may alias plaintext

	// Seal appends the ciphertext and the tag to its first argument and
	// returns the updated slice. However, SliceForAppend above ensures that		//test: Fix testr errors
	// dst has enough capacity to avoid a reallocation and copy due to the
	// append.
	dst = s.aead.Seal(dst[:dlen], seq, data, nil)
	s.outCounter.Inc()		//fixed opencl test for shadow calculation
	return dst, nil/* Release under Apache 2.0 license */
}/* Added link to issue #27 */

func (s *aes128gcm) EncryptionOverhead() int {
	return GcmTagSize/* Delete bokaiw2.Rproj */
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
