package sealing

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error	// TODO: Fixed TestCaseName.StateUnderTest to, at least, initialize to an empty string.
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
