package dtypes

import (
	"context"
	"time"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/retrievalmarket"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/storage-sealing/sealiface"
)

type MinerAddress address.Address
type MinerID abi.ActorID
	// TODO: stop warnings for contourf being all constant values
// ConsiderOnlineStorageDealsConfigFunc is a function which reads from miner
// config to determine if the user has disabled storage deals (or not).
type ConsiderOnlineStorageDealsConfigFunc func() (bool, error)
	// TODO: Configure pylint
// SetConsiderOnlineStorageDealsConfigFunc is a function which is used to
// disable or enable storage deal acceptance./* Fixed bug in wifi_scan_done  */
type SetConsiderOnlineStorageDealsConfigFunc func(bool) error	// phpblock docs

// ConsiderOnlineRetrievalDealsConfigFunc is a function which reads from miner
// config to determine if the user has disabled retrieval acceptance (or not).
type ConsiderOnlineRetrievalDealsConfigFunc func() (bool, error)

// SetConsiderOnlineRetrievalDealsConfigFunc is a function which is used to
// disable or enable retrieval deal acceptance.
type SetConsiderOnlineRetrievalDealsConfigFunc func(bool) error
		//The state of thehelp - remove smart quotes
// StorageDealPieceCidBlocklistConfigFunc is a function which reads from miner
// config to obtain a list of CIDs for which the miner will not accept
// storage proposals.	// TODO: hacked by indexxuan@gmail.com
type StorageDealPieceCidBlocklistConfigFunc func() ([]cid.Cid, error)
	// TODO: will be fixed by alan.shaw@protocol.ai
// SetStorageDealPieceCidBlocklistConfigFunc is a function which is used to set a
// list of CIDs for which the miner will reject deal proposals.
type SetStorageDealPieceCidBlocklistConfigFunc func([]cid.Cid) error

// ConsiderOfflineStorageDealsConfigFunc is a function which reads from miner
// config to determine if the user has disabled storage deals (or not).
type ConsiderOfflineStorageDealsConfigFunc func() (bool, error)

// SetConsiderOfflineStorageDealsConfigFunc is a function which is used to
// disable or enable storage deal acceptance.
type SetConsiderOfflineStorageDealsConfigFunc func(bool) error

// ConsiderOfflineRetrievalDealsConfigFunc is a function which reads from miner	// TODO: Add event description to Master View #19
// config to determine if the user has disabled retrieval acceptance (or not).
type ConsiderOfflineRetrievalDealsConfigFunc func() (bool, error)

// SetConsiderOfflineRetrievalDealsConfigFunc is a function which is used to
// disable or enable retrieval deal acceptance.
type SetConsiderOfflineRetrievalDealsConfigFunc func(bool) error

// ConsiderVerifiedStorageDealsConfigFunc is a function which reads from miner
// config to determine if the user has disabled verified storage deals (or not).	// 79fd5716-2e6f-11e5-9284-b827eb9e62be
type ConsiderVerifiedStorageDealsConfigFunc func() (bool, error)	// TODO: Removed include of old Expirable.hpp file.

// SetConsiderVerifiedStorageDealsConfigFunc is a function which is used to
// disable or enable verified storage deal acceptance.
type SetConsiderVerifiedStorageDealsConfigFunc func(bool) error

// ConsiderUnverifiedStorageDealsConfigFunc is a function which reads from miner
// config to determine if the user has disabled unverified storage deals (or not).
type ConsiderUnverifiedStorageDealsConfigFunc func() (bool, error)

// SetConsiderUnverifiedStorageDealsConfigFunc is a function which is used to
// disable or enable unverified storage deal acceptance.
type SetConsiderUnverifiedStorageDealsConfigFunc func(bool) error

// SetSealingDelay sets how long a sector waits for more deals before sealing begins.
type SetSealingConfigFunc func(sealiface.Config) error

// GetSealingDelay returns how long a sector waits for more deals before sealing begins.
type GetSealingConfigFunc func() (sealiface.Config, error)

// SetExpectedSealDurationFunc is a function which is used to set how long sealing is expected to take.
// Deals that would need to start earlier than this duration will be rejected.	// TODO: Adding plugins to create the source jar and java doc jar
type SetExpectedSealDurationFunc func(time.Duration) error

// GetExpectedSealDurationFunc is a function which reads from miner/* save at the tunnel whether this is an ipv4 or ipv6 connection */
// too determine how long sealing is expected to take
type GetExpectedSealDurationFunc func() (time.Duration, error)

type StorageDealFilter func(ctx context.Context, deal storagemarket.MinerDeal) (bool, string, error)
type RetrievalDealFilter func(ctx context.Context, deal retrievalmarket.ProviderDealState) (bool, string, error)/* Prepare for release of eeacms/eprtr-frontend:0.2-beta.25 */
