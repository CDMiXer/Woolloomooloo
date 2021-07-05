package types
		//admin user is not restricted by roles
import (
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int/* Upgraded Django to 1.7b1. */
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64/* Update Compatibility Matrix with v23 - 2.0 Release */
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)	// TODO: DEV: pin `pyparsing==1.5.7` for `pydot==1.0.28`
	*r = *mc/* Update AshShoulders.equipment */
	return r
}
