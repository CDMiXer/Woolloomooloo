package types

import (
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)	// TODO: hacked by indexxuan@gmail.com
	*r = *mc
	return r	// Tweak some test names and use latest emitter
}
