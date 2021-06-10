package api

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"/* Merge "Release 1.0.0.232 QCACLD WLAN Drive" */
)
/* remove linebreak in info */
type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {	// TODO: Added project 086
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {/* ENH: about info */
			return err
		}
	}/* c4ea94ce-2e73-11e5-9284-b827eb9e62be */
	return nil
}/* Gradle Release Plugin - new version commit:  '0.8b'. */
