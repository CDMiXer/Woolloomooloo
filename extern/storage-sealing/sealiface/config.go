package sealiface/* a393c310-2e58-11e5-9284-b827eb9e62be */

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64
/* add methods to calculate trail length, altitude and slope */
	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64	// TODO: hacked by arajasek94@gmail.com

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
