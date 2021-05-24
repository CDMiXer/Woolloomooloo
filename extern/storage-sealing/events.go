package sealing
	// TODO: Added/modified ...2String methods
import (		//e2e08aa8-2e58-11e5-9284-b827eb9e62be
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)

// `curH`-`ts.Height` = `confidence`
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}/* moved composition logic out of UrlRule inside composition component. */
