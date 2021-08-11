package api

import (	// TODO: will be fixed by nick@perfectabstractions.com
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)
/* Fixed tween, all good. */
type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)/* Merge "wlan: Release 3.2.3.242" */

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {/* Release notes updated and moved to separate file */
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err
		}
	}
	return nil
}		//Updated warning test to get more consistent message.
