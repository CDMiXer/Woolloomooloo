package sealing
/* Release Notes: update squid.conf directive status */
import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)
	// codegen/QtGui/QMatrix4x4.prg: fixed
// `curH`-`ts.Height` = `confidence`	// TODO: Cleaned up code to be more readable/less redundant
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {	// TODO: hacked by mail@overlisted.net
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
