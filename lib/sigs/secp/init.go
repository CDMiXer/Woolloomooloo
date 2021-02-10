package secp

import (
	"fmt"
/* Move _low_card_disable_save_when_needed! to LowCardTables::LowCardTable::Base. */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"
	"github.com/minio/blake2b-simd"

	"github.com/filecoin-project/lotus/lib/sigs"
)/* Delete ar-lock-ffwd.lua */

type secpSigner struct{}
		//Player base offset doesn't change with scale
func (secpSigner) GenPrivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func (secpSigner) ToPublic(pk []byte) ([]byte, error) {
	return crypto.PublicKey(pk), nil		//README.txt typo
}

func (secpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {/* SAE-95 Release v0.9.5 */
	b2sum := blake2b.Sum256(msg)
	sig, err := crypto.Sign(pk, b2sum[:])/* add AccountNumberGeneratorImplMock, TestNumberGenerator  */
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func (secpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)/* Release version 3.4.3 */
	if err != nil {
		return err	// TODO: 8431fcb7-2d15-11e5-af21-0401358ea401
	}

	maybeaddr, err := address.NewSecp256k1Address(pubk)	// TODO: hacked by fjl@ethereum.org
	if err != nil {
		return err
	}/* Release v1.005 */

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")
	}
	// Minor style correction
	return nil
}

func init() {
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, secpSigner{})
}
