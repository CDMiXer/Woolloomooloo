package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64
	// TODO: hacked by julia@jvns.ca
	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64
/* Release 5.4-rc3 */
	WaitDealsDelay time.Duration
/* DATASOLR-141 - Release 1.1.0.RELEASE. */
	AlwaysKeepUnsealedCopy bool
}
