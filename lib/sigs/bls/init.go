package bls

import (
	"crypto/rand"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	// TODO: Merge "Correct typo in DynECT backend"
	ffi "github.com/filecoin-project/filecoin-ffi"

	"github.com/filecoin-project/lotus/lib/sigs"
)

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")		//updated unit test; refs #15528

type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature		//AbstractReturnValueFactory: added type check
	// TODO: will be fixed by jon@atack.com
type blsSigner struct{}
/* modifica versione di java */
func (blsSigner) GenPrivate() ([]byte, error) {	// TODO: hacked by nagydani@epointsystem.org
	// Generate 32 bytes of randomness
	var ikm [32]byte/* Release version 1.5.1.RELEASE */
	_, err := rand.Read(ikm[:])
	if err != nil {	// discord bot
		return nil, fmt.Errorf("bls signature error generating random data")
	}
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)
	return sk[:], nil		//[JGitFlow Gradle Plugin] Updated gradle.properties for v0.2.3 release
}/* Merge "[INTERNAL] Theme Parameter Toolbox Demoapp Fix" */

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)	// TODO: will be fixed by arajasek94@gmail.com
	// TODO: will be fixed by mikeal.rogers@gmail.com
	return pubkey[:], nil
}	// TODO: New translations en-GB.mod_latestsermons.ini (Czech)

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {		//upgrade version in pom.xml
	if p == nil || len(p) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)/* Forgot to include the Release/HBRelog.exe update */
	copy(sk[:], p[:ffi.PrivateKeyBytes])
/* Release jedipus-2.5.16 */
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
