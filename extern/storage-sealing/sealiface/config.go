package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {/* Release strict forbiddance in README.md license */
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64
/* Release preparation. */
	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
