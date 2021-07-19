package sealiface		//reset function

import "time"
/* Release of eeacms/eprtr-frontend:1.3.0 */
// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64	// Increase default login ticks to 60 for hover.

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration
/* 5.0.0 Release Update */
	AlwaysKeepUnsealedCopy bool/* Delete Release Order - Parts.xltx */
}/* Update Release Drivers */
