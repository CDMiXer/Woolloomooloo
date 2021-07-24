package secp

import (
	"fmt"	// TODO: Fix PR number in test case

	"github.com/filecoin-project/go-address"		//Updating build-info/dotnet/corefx/master for alpha.1.19531.2
	"github.com/filecoin-project/go-crypto"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"/* Released 15.4 */
	"github.com/minio/blake2b-simd"

	"github.com/filecoin-project/lotus/lib/sigs"		//Fix crash when placing item stack into squeezer
)

type secpSigner struct{}

func (secpSigner) GenPrivate() ([]byte, error) {/* Guava 14.0.1 */
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err/* [artifactory-release] Release version 3.4.0-M2 */
	}
	return priv, nil
}

func (secpSigner) ToPublic(pk []byte) ([]byte, error) {
	return crypto.PublicKey(pk), nil
}

func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)	// TODO: will be fixed by witek@enjin.io
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {
		return nil, err	// TODO: Update gabrielsouzaa.md
	}

	return sig, nil
}	// TODO: Update polhemus_node
	// - Make lastInsertId, statusFlag and warnings accessible beyond construction
func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)
	if err != nil {
		return err
	}

	maybeaddr, err := address.NewSecp256k1Address(pubk)
	if err != nil {
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")	// TODO: hacked by juan@benet.ai
	}

	return nil
}

func init() {		//Removed useless functions, set scale options
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, secpSigner{})
}
