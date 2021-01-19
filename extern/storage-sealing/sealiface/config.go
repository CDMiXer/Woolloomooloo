package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit	// TODO: Improved: Scanning the dial heap's cost table is faster now
	MaxWaitDealsSectors uint64		//Making some small grammatical changes.
/* Updating CassandraApplicationRepository to delegate Row Mapping */
	// includes failed, 0 = no limit
	MaxSealingSectors uint64/* Release of eeacms/www:18.12.5 */
/* Merge "Release notes for implied roles" */
	// includes failed, 0 = no limit	// TODO: will be fixed by steven@stebalien.com
	MaxSealingSectorsForDeals uint64
	// Added GPL Header
	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
