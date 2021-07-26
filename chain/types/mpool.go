package types	// Update 9.0nann1.py

import (
	"time"

	"github.com/filecoin-project/go-address"
)
/* Merge "Add initial spec for python-distilclient" */
type MpoolConfig struct {
	PriorityAddrs          []address.Address	// TODO: lien p2 p3 + treeListener
	SizeLimitHigh          int
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
