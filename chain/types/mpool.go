package types
	// TODO: will be fixed by davidad@alum.mit.edu
import (
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration		//Delete Perceptron-1.10.py
	GasLimitOverestimation float64
}
/* Fixed cert date */
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r/* 1gpu test script updates for Gtx480  */
}		//Fix bug in Configuration#[]= type checking
