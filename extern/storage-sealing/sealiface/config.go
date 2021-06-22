package sealiface
/*  0.19.4: Maintenance Release (close #60) */
import "time"/* Merge "Release 3.2.3.451 Prima WLAN Driver" */
/* Allow missing possibleNumberPattern when lax */
// this has to be in a separate package to not make lotus API depend on filecoin-ffi		//Add support for GitHub Actions CI build.

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
