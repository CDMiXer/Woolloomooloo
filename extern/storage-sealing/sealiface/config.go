package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64/* Release 0.94.400 */

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}/* Release 3.1.0.M1 */
