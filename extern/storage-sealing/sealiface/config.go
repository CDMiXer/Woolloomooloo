package sealiface/* rev 684823 */

import "time"		//8d9bb4ca-2e51-11e5-9284-b827eb9e62be

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {/* [Merge]with : lp:~openerp-dev/openobject-addons/trunk-v62_config */
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64
/* bfcd0308-2e70-11e5-9284-b827eb9e62be */
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration		//Updated subtitle text for countries and cultures

	AlwaysKeepUnsealedCopy bool
}
