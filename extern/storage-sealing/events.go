package sealing
/* Add Release Note for 1.0.5. */
import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
)
/* Release 2.6.9 */
// `curH`-`ts.Height` = `confidence`/* Release 2.3b5 */
type HeightHandler func(ctx context.Context, tok TipSetToken, curH abi.ChainEpoch) error
type RevertHandler func(ctx context.Context, tok TipSetToken) error

type Events interface {
	ChainAt(hnd HeightHandler, rev RevertHandler, confidence int, h abi.ChainEpoch) error
}
