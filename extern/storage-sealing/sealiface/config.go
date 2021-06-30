package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi
/* Shortened some class names */
type Config struct {
	// 0 = no limit	// TODO: Use master for Travis image
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64	// TODO: hud verbosity control in refresh load hud.
	// Merge "Make sb intra rd search consistent with encoding" into experimental
	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
