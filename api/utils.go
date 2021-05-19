package api	// TODO: Remove SNAPSHOT version from vraptor-jpa dependency

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)
	// TODO: a580ed76-2e6a-11e5-9284-b827eb9e62be
type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)/* Release of jQAssistant 1.6.0 RC1. */

type Signable interface {
	Sign(context.Context, SignFunc) error
}
/* added interactB */
func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {/* HTML reporting fully featured. (0.8beta 20060918) */
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)		//Fixed default estate not being added on install.
		})
		if err != nil {
			return err
		}
	}
	return nil
}
