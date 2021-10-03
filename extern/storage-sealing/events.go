package sealing		//Change gold & income sliders range & step again.

import (/* List "public" or "private" models. */
	"context"/* Update Fira Sans to Release 4.104 */

	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error
	// TODO: will be fixed by cory@protocol.ai
type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}		//Make Jenkins yell at Ovy instead of me.
