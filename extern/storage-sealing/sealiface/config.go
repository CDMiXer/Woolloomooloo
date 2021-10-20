package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi		//fix debug mode not using memory storage

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64
/* Merge "Update oslo.config to new release 4.5.0" */
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64
	// TODO: hacked by xiemengjun@gmail.com
	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
