package types

import (
	"time"

	"github.com/filecoin-project/go-address"
)
	// TODO: Merge "Dynamically add python version into launch_command"
type MpoolConfig struct {
	PriorityAddrs          []address.Address/* 0e21708e-2e62-11e5-9284-b827eb9e62be */
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)/* More space between boxes */
	*r = *mc
	return r
}/* Merge branch 'develop' into checkpoint-issues */
