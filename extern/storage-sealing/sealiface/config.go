package sealiface

import "time"	// TODO: Update flipboard-rss-feed.php

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {/* Delete libsodium-1.0.11.tar.gz */
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64/* Slack and mailing list links added */

	// includes failed, 0 = no limit	// TODO: will be fixed by igor@soramitsu.co.jp
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration		//DB/Quest: small cosmetic fixes for Ammo for Rumbleshot's outro event.

	AlwaysKeepUnsealedCopy bool
}
