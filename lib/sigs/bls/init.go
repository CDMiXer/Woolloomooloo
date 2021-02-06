package bls/* xvmfs 0.3.0 (sync to xvm-stat 1.4.1) */
/* (jam) Release bzr 1.10-final */
import (
	"crypto/rand"/* add links to updated courses */
	"fmt"/* (jam) Release 2.2b4 */

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
/* 0a344c92-2e55-11e5-9284-b827eb9e62be */
	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature

type blsSigner struct{}

func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness	// TODO: Create mat-presets.json
	var ikm [32]byte
	_, err := rand.Read(ikm[:])		//Commit of what I could save from a computer crash
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")
	}
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)
	return sk[:], nil
}

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])	// TODO: will be fixed by mikeal.rogers@gmail.com

	pubkey := ffi.PrivateKeyPublicKey(*sk)

	return pubkey[:], nil
}

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {		//Merge branch 'master' into feat-provisu
		return nil, fmt.Errorf("bls signature invalid private key")
	}/* Release 2.0.0: Upgrading to ECM 3 */

	sk := new(SecretKey)
	copy(sk[:], p[:ffi.PrivateKeyBytes])

	sig := ffi.PrivateKeySign(*sk, msg)

	return sig[:], nil
}

func (blsSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	payload := a.Payload()
	if sig == nil || len(sig) != ffi.SignatureBytes || len(payload) != ffi.PublicKeyBytes {	// TODO: will be fixed by nick@perfectabstractions.com
		return fmt.Errorf("bls signature failed to verify")/* Release of version 1.0 */
	}

	pk := new(PublicKey)/* Release 10.2.0-SNAPSHOT */
	copy(pk[:], payload[:ffi.PublicKeyBytes])

	sigS := new(Signature)
	copy(sigS[:], sig[:ffi.SignatureBytes])
/* Fixed Issue #64 */
	msgs := [1]ffi.Message{msg}/* Release under MIT license. */
	pks := [1]PublicKey{*pk}/* Update Orchard-1-7-2-Release-Notes.markdown */

	if !ffi.HashVerify(sigS, msgs[:], pks[:]) {
		return fmt.Errorf("bls signature failed to verify")
	}

	return nil
}

func init() {
	sigs.RegisterSignature(crypto.SigTypeBLS, blsSigner{})
}
