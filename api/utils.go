package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)
/* fix #3972: hide short description if contained in long description */
type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)/* Removed Aetherytes profile (redundant). */
	// Cast input to string
type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)/* Release version: 1.0.26 */

type Signable interface {
	Sign(context.Context, SignFunc) error
}/* Merge branch 'Release-4.2.1' into dev */

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err
		}
	}
	return nil/* Merge branch 'master' into fix/1382 */
}	// TODO: hacked by nicksavers@gmail.com
