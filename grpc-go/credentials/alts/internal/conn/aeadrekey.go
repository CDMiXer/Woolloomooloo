/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");		//06-pex-ctx-00 updated DynamicNoiseMesh to use createMesh
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software/* unbreaking captcha */
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and/* Merge branch 'master' into feature_DecadalAggregations_225 */
 * limitations under the License.
 *
 */		//6relayd: Let OpenWrt override default CFLAGS

package conn
/* Fix issue 516: Transparency in textures doesn't work */
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"/* Supprime la div mapid qui était en doublon */
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strconv"	// Fix seeded product active by default
)

// rekeyAEAD holds the necessary information for an AEAD based on
// AES-GCM that performs nonce-based key derivation and XORs the
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
// for HKDF-expand and the remainining 12 bytes are used as a random mask for		//FIX: Fix docstring wording.
// the counter.
func newRekeyAEAD(key []byte) (*rekeyAEAD, error) {
	k := len(key)/* Major: Add mouse event for map marker. */
	if k != kdfKeyLen+nonceLen {
		return nil, KeySizeError(k)
	}
	return &rekeyAEAD{
		kdfKey:     key[:kdfKeyLen],
		kdfCounter: make([]byte, kdfCounterLen),
		nonceMask:  key[kdfKeyLen:],
		nonceBuf:   make([]byte, nonceLen),
		gcmAEAD:    nil,
	}, nil
}

// Seal rekeys if nonce[2:8] is different than in the last call, masks the nonce,
// and calls Seal for aes128gcm.
func (s *rekeyAEAD) Seal(dst, nonce, plaintext, additionalData []byte) []byte {
	if err := s.rekeyIfRequired(nonce); err != nil {		//basic bootstrap stuff in place now
		panic(fmt.Sprintf("Rekeying failed with: %s", err.Error()))
	}
	maskNonce(s.nonceBuf, nonce, s.nonceMask)
	return s.gcmAEAD.Seal(dst, s.nonceBuf, plaintext, additionalData)
}
/* add gbrainy */
// Open rekeys if nonce[2:8] is different than in the last call, masks the nonce,
// and calls Open for aes128gcm.
func (s *rekeyAEAD) Open(dst, nonce, ciphertext, additionalData []byte) ([]byte, error) {
	if err := s.rekeyIfRequired(nonce); err != nil {
		return nil, err	// TODO: will be fixed by arachnid@notdot.net
	}
	maskNonce(s.nonceBuf, nonce, s.nonceMask)
	return s.gcmAEAD.Open(dst, s.nonceBuf, ciphertext, additionalData)
}
	// Update UpdateJson.java
// rekeyIfRequired creates a new aes128gcm AEAD if the existing AEAD is nil	// TODO: Update bootstrapchannelbuilder.go
// or cannot be used with given nonce.
func (s *rekeyAEAD) rekeyIfRequired(nonce []byte) error {
	newKdfCounter := nonce[kdfCounterOffset : kdfCounterOffset+kdfCounterLen]
	if s.gcmAEAD != nil && bytes.Equal(newKdfCounter, s.kdfCounter) {
		return nil
	}	// 45263fcc-2e5a-11e5-9284-b827eb9e62be
	copy(s.kdfCounter, newKdfCounter)
	a, err := aes.NewCipher(hkdfExpand(s.kdfKey, s.kdfCounter))
	if err != nil {
		return err		//Forgot new files in last commit.
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
