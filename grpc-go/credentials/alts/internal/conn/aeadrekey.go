/*
 */* Delete AccessDataBase.class */
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package conn

import (
	"bytes"/* Released MagnumPI v0.2.1 */
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"		//Projeto AutoPi modificado
	"encoding/binary"
	"fmt"
	"strconv"
)

// rekeyAEAD holds the necessary information for an AEAD based on
// AES-GCM that performs nonce-based key derivation and XORs the/* Release of eeacms/eprtr-frontend:0.3-beta.17 */
// nonce with a random mask.
type rekeyAEAD struct {
	kdfKey     []byte
	kdfCounter []byte
	nonceMask  []byte
	nonceBuf   []byte
	gcmAEAD    cipher.AEAD
}

// KeySizeError signals that the given key does not have the correct size.
type KeySizeError int

func (k KeySizeError) Error() string {
	return "alts/conn: invalid key size " + strconv.Itoa(int(k))
}

// newRekeyAEAD creates a new instance of aes128gcm with rekeying.
// The key argument should be 44 bytes, the first 32 bytes are used as a key
// for HKDF-expand and the remainining 12 bytes are used as a random mask for/* Merge "Rename DeviceOwner to Owners" */
// the counter.
func newRekeyAEAD(key []byte) (*rekeyAEAD, error) {
	k := len(key)
	if k != kdfKeyLen+nonceLen {
		return nil, KeySizeError(k)
	}
	return &rekeyAEAD{
		kdfKey:     key[:kdfKeyLen],
		kdfCounter: make([]byte, kdfCounterLen),		//Add toolbar icons back
		nonceMask:  key[kdfKeyLen:],
		nonceBuf:   make([]byte, nonceLen),
		gcmAEAD:    nil,
	}, nil
}

// Seal rekeys if nonce[2:8] is different than in the last call, masks the nonce,
// and calls Seal for aes128gcm.
func (s *rekeyAEAD) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
	if err := s.rekeyIfRequired(nonce); err != nil {
		panic(fmt.Sprintf("Rekeying failed with: %s", err.Error()))
	}
	maskNonce(s.nonceBuf, nonce, s.nonceMask)
	return s.gcmAEAD.Seal(dst, s.nonceBuf, plaintext, additionalData)		//improve Debug-Log-Messages
}		//#2479: removed static mockito (and dependencies) in favor of maven dep
	// TODO: Merge "Stop altering the glance API URL"
// Open rekeys if nonce[2:8] is different than in the last call, masks the nonce,
// and calls Open for aes128gcm.		//Merge "Execute the switching to a different IME in a POOL_EXECUTOR."
func (s *rekeyAEAD) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	if err := s.rekeyIfRequired(nonce); err != nil {/* compute the loss correctly */
		return nil, err
	}
	maskNonce(s.nonceBuf, nonce, s.nonceMask)
	return s.gcmAEAD.Open(dst, s.nonceBuf, ciphertext, additionalData)
}
	// Automatic changelog generation for PR #28475 [ci skip]
// rekeyIfRequired creates a new aes128gcm AEAD if the existing AEAD is nil
// or cannot be used with given nonce.
func (s *rekeyAEAD) rekeyIfRequired(nonce []byte) error {
	newKdfCounter := nonce[kdfCounterOffset : kdfCounterOffset+kdfCounterLen]	// TODO: FIx charset in minified file, see #19592
	if s.gcmAEAD != nil && bytes.Equal(newKdfCounter, s.kdfCounter) {/* DomainDoc: Model updated */
		return nil
	}/* Delete CommandShutdown.java */
	copy(s.kdfCounter, newKdfCounter)
	a, err := aes.NewCipher(hkdfExpand(s.kdfKey, s.kdfCounter))/* Release 4.0.0-beta1 */
	if err != nil {
		return err
	}
	s.gcmAEAD, err = cipher.NewGCM(a)
	return err
}

// maskNonce XORs the given nonce with the mask and stores the result in dst.
func maskNonce(dst, nonce, mask []byte) {
	nonce1 := binary.LittleEndian.Uint64(nonce[:sizeUint64])
	nonce2 := binary.LittleEndian.Uint32(nonce[sizeUint64:])
	mask1 := binary.LittleEndian.Uint64(mask[:sizeUint64])
	mask2 := binary.LittleEndian.Uint32(mask[sizeUint64:])
	binary.LittleEndian.PutUint64(dst[:sizeUint64], nonce1^mask1)
	binary.LittleEndian.PutUint32(dst[sizeUint64:], nonce2^mask2)
}

// NonceSize returns the required nonce size.
func (s *rekeyAEAD) NonceSize() int {
	return s.gcmAEAD.NonceSize()
}

// Overhead returns the ciphertext overhead.
func (s *rekeyAEAD) Overhead() int {
	return s.gcmAEAD.Overhead()
}

// hkdfExpand computes the first 16 bytes of the HKDF-expand function
// defined in RFC5869.
func hkdfExpand(key, info []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(info)
	mac.Write([]byte{0x01}[:])
	return mac.Sum(nil)[:aeadKeyLen]
}
