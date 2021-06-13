package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64/* Release: 5.6.0 changelog */
		//add a new step to the legacy migration checklist
	// includes failed, 0 = no limit
	MaxSealingSectors uint64/* Merge "Support both VLAN and VXLAN on a single cluster" */
/* Italian locale v.2.3 added */
	// includes failed, 0 = no limit		//Added new test and simple classes for marker data
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
