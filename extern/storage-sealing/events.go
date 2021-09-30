package sealing

import (	// TODO: Merge "esoc: Add debug engine for external modems."
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)
		//Fixed ref counting bug for loading templates
// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}		//CF/BF - delete some unused code from BST.
