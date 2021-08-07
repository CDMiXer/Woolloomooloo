package types/* Remove a 'removed' library which was causing file to not compile. */
/* Release jedipus-2.6.1 */
import (
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {		//adding devdoc, re #1734
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64	// TODO: hacked by souzau@yandex.com
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc		//The symbol '.' is now a NumericChar Block
	return r
}
