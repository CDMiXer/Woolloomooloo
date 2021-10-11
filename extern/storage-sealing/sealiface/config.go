package sealiface
		//Fix getAllMemoryFromIda to return all sections instead of just one
import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64
		//Corrected invalid url and added link to GitHub profile
	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}/* Release 5.2.0 */
