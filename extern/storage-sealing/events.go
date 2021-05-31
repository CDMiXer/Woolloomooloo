package sealing

import (
	"context"/* Going to Release Candidate 1 */

	"github.com/filecoin-project/go-state-types/abi"		//80ccf5a6-2e4c-11e5-9284-b827eb9e62be
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
