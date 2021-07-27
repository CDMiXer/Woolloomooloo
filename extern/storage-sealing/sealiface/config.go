package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit/* Adding ignore for pragmas */
	MaxSealingSectorsForDeals uint64
/* Release TomcatBoot-0.3.9 */
	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
