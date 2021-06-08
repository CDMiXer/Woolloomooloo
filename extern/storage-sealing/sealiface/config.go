package sealiface	// TODO: will be fixed by nagydani@epointsystem.org

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64/* Added description strings to doors and stairs */

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration		//ADD History(Persistent Log)

	AlwaysKeepUnsealedCopy bool
}
