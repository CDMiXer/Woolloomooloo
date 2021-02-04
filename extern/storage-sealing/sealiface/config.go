package sealiface

import "time"
/* Merge pull request #94 from atdrago/feature/adam/travis-setup */
// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit/* Merge branch 'master' into components */
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64
/* Make-Release */
	WaitDealsDelay time.Duration
		//Delete assign33.t
	AlwaysKeepUnsealedCopy bool
}
