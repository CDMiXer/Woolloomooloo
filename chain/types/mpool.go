package types

import (	// added readme.md file
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64/* Bug 3941: Release notes typo */
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}
		//F5 - update
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}
