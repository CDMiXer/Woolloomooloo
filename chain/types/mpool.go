package types
	// TODO: will be fixed by steven@stebalien.com
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
	GasLimitOverestimation float64		//Use arrow functions
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)/* Delete Release.md */
	*r = *mc		//added Experiment.getExperimentByName
	return r
}
