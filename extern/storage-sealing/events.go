package sealing

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)
	// TODO: will be fixed by joshua@yottadb.com
// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error/* Entwicklungs-Bildressourcen hinzugefügt */
}
