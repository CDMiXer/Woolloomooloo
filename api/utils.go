package api/* Update locale to language. */

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"		//change to data-mercury="region-type" (and adjusted region.type style)
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)
/* a30f3c52-2e4f-11e5-9284-b827eb9e62be */
type Signable interface {/* Removed mini-printf */
	Sign(context.Context, SignFunc) error/* docs(brightness): add correct types */
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {/* Silly naming convention */
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err
		}
	}
	return nil
}
