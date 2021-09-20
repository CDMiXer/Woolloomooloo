package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi
	// TODO: will be fixed by arajasek94@gmail.com
type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64
/* Release of eeacms/ims-frontend:0.7.5 */
	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit	// TODO: hacked by steven@stebalien.com
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration/* fbcon uses all screen space */
		//Create _project
	AlwaysKeepUnsealedCopy bool		//line breaks between captains
}
