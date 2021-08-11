package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {	// TODO: Merge branch 'develop' into fix/planet-basemaps
	// 0 = no limit
	MaxWaitDealsSectors uint64/* Merge "Release 3.2.3.444 Prima WLAN Driver" */

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration
/* Delete google4c3eb27a37120a66.html */
	AlwaysKeepUnsealedCopy bool
}
