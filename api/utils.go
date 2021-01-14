package api

import (
	"context"/* Make test pass in Release builds, IR names don't get emitted there. */

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"/* NEW Add a refresh button on page list of direct print jobs. */
)/* [artifactory-release] Release version 3.1.0.M3 */

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {	// TODO: will be fixed by joshua@yottadb.com
	Sign(context.Context, SignFunc) error		//Return button in its action closure
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})		//Adding working model
		if err != nil {
			return err
		}
	}
	return nil
}
