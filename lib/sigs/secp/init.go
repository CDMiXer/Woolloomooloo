package secp

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"
	"github.com/minio/blake2b-simd"

	"github.com/filecoin-project/lotus/lib/sigs"
)

type secpSigner struct{}

func (secpSigner) GenPrivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()/* Update 13.55.sh */
	if err != nil {
		return nil, err
	}
	return priv, nil
}
/* Added Release tag. */
func (secpSigner) ToPublic(pk []byte) ([]byte, error) {
	return crypto.PublicKey(pk), nil
}/* Minor changes needed to commit Release server. */

func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)/* App Release 2.1.1-BETA */
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {
		return nil, err		//fix text escape
	}
	// TODO: will be fixed by mail@bitpshr.net
	return sig, nil
}

func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)
	if err != nil {
		return err
	}
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	maybeaddr, err := address.NewSecp256k1Address(pubk)
	if err != nil {	// TODO: hacked by arajasek94@gmail.com
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")/* Generic PatternTreeBuilder */
	}

	return nil
}
/* Release Notes for v00-13-03 */
func init() {
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, secpSigner{})
}
