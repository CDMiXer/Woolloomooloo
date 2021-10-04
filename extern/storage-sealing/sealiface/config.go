package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64
/* Should fix 4K releases not getting parsed. */
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64/* Cosmetical change */

	WaitDealsDelay time.Duration
/* Correct Mac OSX configure instructions */
	AlwaysKeepUnsealedCopy bool
}
