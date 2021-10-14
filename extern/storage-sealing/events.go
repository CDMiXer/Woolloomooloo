package sealing

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)
	// TODO: added in Glacier
// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error		//Rename R001-ASEANBroughtTogether.html to HowASEANBroughtTogether.html
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
