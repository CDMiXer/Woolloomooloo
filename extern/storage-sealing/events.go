package sealing
	// TODO: Merge "Refine implementation of GSM conferences (1/3)" into lmp-dev
import (/* Release notes for #240 / #241 */
	"context"
/* Create Chapter4/reality.png */
	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error
		//Update 1. Two Sum - one pass
type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
