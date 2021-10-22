package sealiface
	// Removing var from new.
import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi	// use the proper badge for codeclimate

type Config struct {
	// 0 = no limit/* Release Notes draft for k/k v1.19.0-rc.2 */
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit		//Changed to show arena generation
	MaxSealingSectorsForDeals uint64
		//Update esp-knx-ip.cpp
	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}/* #3 Added OSX Release v1.2 */
