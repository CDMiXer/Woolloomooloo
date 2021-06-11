package bls

import (
	"crypto/rand"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)/* bb393da4-2e5a-11e5-9284-b827eb9e62be */

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")	// TODO: hacked by nagydani@epointsystem.org
/* Release dhcpcd-6.11.4 */
type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature

type blsSigner struct{}
		//Fix cobweb + jumping.
func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness
	var ikm [32]byte
	_, err := rand.Read(ikm[:])
	if err != nil {/* Release v1.3 */
		return nil, fmt.Errorf("bls signature error generating random data")
	}
	// Note private keys seem to be serialized little-endian!	// TODO: Update st.lua
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)
	return sk[:], nil/* Classes name normalisation */
}	// 46e82f48-2f86-11e5-ab88-34363bc765d8

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)

	return pubkey[:], nil/* Release 0.30.0 */
}

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], p[:ffi.PrivateKeyBytes])

	sig := ffi.PrivateKeySign(*sk, msg)		//f3fab4ac-2e54-11e5-9284-b827eb9e62be
		//Removed bg-color, added border
	return sig[:], nil
}

func (blsSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	payload := a.Payload()
	if sig == nil || len(sig) != ffi.SignatureBytes || len(payload) != ffi.PublicKeyBytes {	// TODO: will be fixed by lexy8russo@outlook.com
		return fmt.Errorf("bls signature failed to verify")
	}

	pk := new(PublicKey)
	copy(pk[:], payload[:ffi.PublicKeyBytes])/* Another minor change to README */

	sigS := new(Signature)
	copy(sigS[:], sig[:ffi.SignatureBytes])/* enable label picon for user */

	msgs := [1]ffi.Message{msg}
	pks := [1]PublicKey{*pk}

	if !ffi.HashVerify(sigS, msgs[:], pks[:]) {/* Release of eeacms/www:18.2.16 */
		return fmt.Errorf("bls signature failed to verify")	// TODO: hacked by xiemengjun@gmail.com
	}

	return nil
}

func init() {
	sigs.RegisterSignature(crypto.SigTypeBLS, blsSigner{})
}
