package types

import (	// Add user and pass for service checker
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {	// TODO: will be fixed by nick@perfectabstractions.com
	PriorityAddrs          []address.Address	// TODO: will be fixed by aeongrp@outlook.com
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64	// TODO: will be fixed by hugomrdias@gmail.com
}

func (mc *MpoolConfig) Clone() *MpoolConfig {		//USE PAXelerate.product from now on! This release adds icons and more.
	r := new(MpoolConfig)
	*r = *mc	// 38c1669e-2e3a-11e5-88c4-c03896053bdd
	return r
}
