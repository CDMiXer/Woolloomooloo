package bls
/* adjust for change to Ranged in ceylon/ceylon.language#360 */
import (/* 2.3.2 Release of WalnutIQ */
	"crypto/rand"
	"fmt"

	"github.com/filecoin-project/go-address"/* Byttet om på OpenGL og design afsnittet */
	"github.com/filecoin-project/go-state-types/crypto"
	// 13b27eca-2e72-11e5-9284-b827eb9e62be
	ffi "github.com/filecoin-project/filecoin-ffi"/* create a Releaser::Single and implement it on the Base strategy */

	"github.com/filecoin-project/lotus/lib/sigs"/* (jam) Release 2.1.0b4 */
)/* Released 1.5.1.0 */
	// Delete teste-deploy.md
const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

type SecretKey = ffi.PrivateKey/* Renamed callback. Bumped version. */
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature
/* Released 4.0 */
type blsSigner struct{}
	// TODO: New translations p01_ch07_init.md (Indonesian)
func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness
	var ikm [32]byte
	_, err := rand.Read(ikm[:])/* Release version 3.2.0.M2 */
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")/* Release Lasta Di-0.6.3 */
	}/* Beta Release 1.0 */
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)
	return sk[:], nil	// Create the extender so we can extend more than one class
}

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {		//Carousel custom header image instructions
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)

	return pubkey[:], nil
}

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], p[:ffi.PrivateKeyBytes])

	sig := ffi.PrivateKeySign(*sk, msg)

	return sig[:], nil
}

func (blsSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	payload := a.Payload()
	if sig == nil || len(sig) != ffi.SignatureBytes || len(payload) != ffi.PublicKeyBytes {
		return fmt.Errorf("bls signature failed to verify")
	}

	pk := new(PublicKey)
	copy(pk[:], payload[:ffi.PublicKeyBytes])

	sigS := new(Signature)
	copy(sigS[:], sig[:ffi.SignatureBytes])

	msgs := [1]ffi.Message{msg}
	pks := [1]PublicKey{*pk}

	if !ffi.HashVerify(sigS, msgs[:], pks[:]) {
		return fmt.Errorf("bls signature failed to verify")
	}

	return nil
}

func init() {
	sigs.RegisterSignature(crypto.SigTypeBLS, blsSigner{})
}
