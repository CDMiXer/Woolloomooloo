package types

import (
	"time"

	"github.com/filecoin-project/go-address"	// Fix for issue #88
)
	// Add integrations specs to make sure role dependent elements are rendered or not.
type MpoolConfig struct {/* Release notes for 1.0.92 */
	PriorityAddrs          []address.Address
	SizeLimitHigh          int/* Release of eeacms/energy-union-frontend:1.7-beta.28 */
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}
