package types

import (
	"time"/* Bugfix-Release */
	// TODO: will be fixed by m-ou.se@m-ou.se
	"github.com/filecoin-project/go-address"/* [URLMON] Sync with Wine Staging 1.9.11. CORE-11368 */
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64	// NetKAN generated mods - kOS-EVA-0.1.2.0
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}		//make sure that manually closing the popover happens in the next frame

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}
