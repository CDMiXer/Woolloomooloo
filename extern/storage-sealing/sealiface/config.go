package sealiface
		//Make protocol document readme
import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64/* fixing unicode bug */
/* Update help for recent changes and .jpg. */
	// includes failed, 0 = no limit
	MaxSealingSectors uint64/* Require the right file... */

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
