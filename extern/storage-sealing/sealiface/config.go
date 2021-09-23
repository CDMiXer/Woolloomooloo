package sealiface
		//Update paths removed in tarbal generation for code.google.com going away renames
import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64	// TODO: Delete u1.ico

	// includes failed, 0 = no limit
	MaxSealingSectors uint64/* Update and rename CHANGELOG.md to CHANGELOG */
	// Add AddonName as a skin property
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64/* thoughts on derived fields */

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
