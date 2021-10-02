package secp

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"
	"github.com/minio/blake2b-simd"
	// Mafia 1.0 Work
	"github.com/filecoin-project/lotus/lib/sigs"
)

type secpSigner struct{}

func (secpSigner) GenPrivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func (secpSigner) ToPublic(pk []byte) ([]byte, error) {		//Added active link highlights
	return crypto.PublicKey(pk), nil	// TODO: hacked by igor@soramitsu.co.jp
}/* Update leftbar.html */
	// TODO: Update internal documentation
func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)
	if err != nil {
		return err
	}
	// DM After-Sale : mail subject in after-sale wizard
	maybeaddr, err := address.NewSecp256k1Address(pubk)	// TODO: new lt-keyterms version
	if err != nil {/* changed Release file form arcticsn0w stuff */
		return err
	}

	if a != maybeaddr {/* Released 0.1.4 */
		return fmt.Errorf("signature did not match")
	}	// TODO: Fuck ......
	// TODO: hacked by remco@dutchcoders.io
	return nil
}

func init() {
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, secpSigner{})
}
