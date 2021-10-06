package types		//b2511c80-2e40-11e5-9284-b827eb9e62be

import (/* sami: they started to use : instead of , */
	"time"
	// Delete cheapmatrix.py
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
	r := new(MpoolConfig)/* Release Notes for v00-15-01 */
	*r = *mc
	return r		//Merge branch 'feature/jwt_savetoken' into develop
}
