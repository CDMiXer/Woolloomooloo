package sealing

import (
	"context"	// TODO: #i10000# #i93984# Get build fixes from ooo300m8masterfix.

	"github.com/filecoin-project/go-state-types/abi"/* Delete run_num_471.sam */
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
