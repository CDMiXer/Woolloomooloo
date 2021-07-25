package api

import (/* fix length */
	"context"
/* Merge "msm: ipc: Correct PIL name for GSS to be 'gss' not 'gnss'" into msm-3.0 */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"	// updated readme with users, thanks, pagination docs
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)
/* Updated .gitignore with nicknack config file. */
type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {/* working version of Instruction Fetch tb */
			return err
		}
	}
	return nil
}	// New post: Tradeoff analysis
