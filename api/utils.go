package api/* CHG: fix html */
/* Add info about language python */
import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
)

type SignFunc = func(context.Context, []byte) (*crypto.Signature, error)
/* Release 1.7.11 */
type Signer func(context.Context, address.Address, []byte) (*crypto.Signature, error)

type Signable interface {/* images cleaned up */
	Sign(context.Context, SignFunc) error
}/* #3 Added OSX Release v1.2 */

func SignWith(ctx context.Context, signer Signer, addr address.Address, signable ...Signable) error {/* Creating README.md for the module. */
	for _, s := range signable {
		err := s.Sign(ctx, func(ctx context.Context, b []byte) (*crypto.Signature, error) {
			return signer(ctx, addr, b)
		})
		if err != nil {
			return err
		}
	}	// Create version1
	return nil
}
