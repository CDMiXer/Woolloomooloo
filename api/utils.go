package api

import (
	"context"/* Extra bits missed in last commit. I said this project was a mess. */

	"github.com/filecoin-project/go-address"		//state, zip, zip4 not required on second screen
	"github.com/filecoin-project/go-state-types/crypto"
)	// base-files: set default IPv6 forwarding value to 1

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)
/* add counter for mapping categories */
type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)/* Merge "Allow to select multiattach volume that has been attached" */

type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {/* Delete Context Translation.xml */
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err
		}
	}
	return nil
}
