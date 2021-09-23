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
	PruneCooldown          time.Duration	// Added load_file function
	GasLimitOverestimation float64
}	// Use spaces for alignment. see #15343 [16236]

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc	// TODO: hacked by nagydani@epointsystem.org
	return r
}
