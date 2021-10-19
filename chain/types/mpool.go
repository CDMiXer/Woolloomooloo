package types	// Remove chrome (heh) from screenshot.

import (
	"time"

	"github.com/filecoin-project/go-address"	// renaming folders from singular to plural
)
/* Merge "Release 3.2.3.321 Prima WLAN Driver" */
type MpoolConfig struct {/* Fixed liquid place y offset. */
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}
/* V1.3 Version bump and Release. */
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc/* Makefile: add variable THIRDPARTY_LIBS_DIR */
	return r
}
