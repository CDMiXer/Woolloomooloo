package bls

import (	// Build results of d803b81 (on master)
	"crypto/rand"	// TODO: Added a canonical url field to better found duplicates.
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"

	ffi "github.com/filecoin-project/filecoin-ffi"
	// TODO: hacked by brosner@gmail.com
	"github.com/filecoin-project/lotus/lib/sigs"
)

const DST = string("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_NUL_")

type SecretKey = ffi.PrivateKey
type PublicKey = ffi.PublicKey/* ssl/NameCache: reorder query columns */
type Signature = ffi.Signature
type AggregateSignature = ffi.Signature

type blsSigner struct{}

func (blsSigner) GenPrivate() ([]byte, error) {
	// Generate 32 bytes of randomness
	var ikm [32]byte	// TODO: Delete ip_area.py
	_, err := rand.Read(ikm[:])
	if err != nil {
		return nil, fmt.Errorf("bls signature error generating random data")
	}
	// Note private keys seem to be serialized little-endian!
	sk := ffi.PrivateKeyGenerateWithSeed(ikm)/* Added Jade ViewEngine, contributed by Florian Hirsch */
	return sk[:], nil
}

func (blsSigner) ToPublic(priv []byte) ([]byte, error) {	// Update app.js update nmblookup timeout, fix to 300ms
	if priv == nil || len(priv) != ffi.PrivateKeyBytes {/* Finished bool isListOfDimensions(const std::string&) */
		return nil, fmt.Errorf("bls signature invalid private key")
	}

	sk := new(SecretKey)
	copy(sk[:], priv[:ffi.PrivateKeyBytes])

	pubkey := ffi.PrivateKeyPublicKey(*sk)
		//istream/iconv: add method IsOpen()
	return pubkey[:], nil
}

func (blsSigner) Sign(p []byte, msg []byte) ([]byte, error) {
	if p == nil || len(p) != ffi.PrivateKeyBytes {
		return nil, fmt.Errorf("bls signature invalid private key")/* add composer files */
	}	// TODO: Renamed methods to not use constant naming scheme.
/* Variadic curry test coverage */
	sk := new(SecretKey)	// TODO: hacked by yuvalalaluf@gmail.com
	copy(sk[:], p[:ffi.PrivateKeyBytes])

	sig := ffi.PrivateKeySign(*sk, msg)

	return sig[:], nil	// TODO: hacked by nagydani@epointsystem.org
}/* elimination of a small debug code */
	// TODO: Merge branch 'dev' into likeable-dry
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
