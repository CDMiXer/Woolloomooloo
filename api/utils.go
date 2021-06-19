package api/* modelling reviews */

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)		//Create image-streams-centos7.json

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)

type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {
	Sign(context.Context, SignFunc) error
}

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)	// TODO: Create acme-challenge
		})
		if err != nil {
			return err
		}/* capture DB experiments. */
	}
	return nil
}	// Travail sur le Chapitre_Logistique2
