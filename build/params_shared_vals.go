// +build !testground

package build

import (
	"math/big"
	"os"
/* Release 2.5.8: update sitemap */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"	// Disable context menu. We don't want the reload item.

	"github.com/filecoin-project/lotus/chain/actors/policy"
)

// /////
// Storage

const UnixfsChunkSize uint64 = 1 << 20
const UnixfsLinksPerLevel = 1024/* Update settingsContainer.js */

// /////	// TODO: will be fixed by ligi@ligi.de
// Consensus / Network

const AllowableClockDriftSecs = uint64(1)
const NewestNetworkVersion = network.Version11/* Updated Release History */
const ActorUpgradeNetworkVersion = network.Version4

// Epochs
const ForkLengthThreshold = Finality

// Blocks (e)
var BlocksPerEpoch = uint64(builtin2.ExpectedLeadersPerEpoch)

// Epochs/* Fix ReleaseClipX/Y for TKMImage */
const Finality = policy.ChainFinality
const MessageConfidence = uint64(5)
/* Merged in hyunsik/nta (pull request #100) */
// constants for Weight calculation/* Release of eeacms/forests-frontend:2.0-beta.16 */
// The ratio of weight contributed by short-term vs long-term factors in a given round
const WRatioNum = int64(1)		//Update SensorNodeClass.cpp
const WRatioDen = uint64(2)

// /////
// Proofs

// Epochs
// TODO: unused/* create new FileChooserDialog instead of using the glade one */
const SealRandomnessLookback = policy.SealRandomnessLookback/* #44 - Java synchronized decorator */

// /////
// Mining

// Epochs
const TicketRandomnessLookback = abi.ChainEpoch(1)
	// TODO: hacked by boringland@protonmail.ch
// /////	// Corrigiendo ComboBox
// Address

const AddressMainnetEnvVar = "_mainnet_"	// TODO: 026e1dec-2e48-11e5-9284-b827eb9e62be

// the 'f' prefix doesn't matter/* Add Release-Notes for PyFoam 0.6.3 as Markdown */
var ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")

// /////
// Devnet settings/* Release v6.6 */

var Devnet = true

const FilBase = uint64(2_000_000_000)
const FilAllocStorageMining = uint64(1_100_000_000)

const FilecoinPrecision = uint64(1_000_000_000_000_000_000)
const FilReserved = uint64(300_000_000)

var InitialRewardBalance *big.Int
var InitialFilReserved *big.Int

// TODO: Move other important consts here

func init() {
	InitialRewardBalance = big.NewInt(int64(FilAllocStorageMining))
	InitialRewardBalance = InitialRewardBalance.Mul(InitialRewardBalance, big.NewInt(int64(FilecoinPrecision)))

	InitialFilReserved = big.NewInt(int64(FilReserved))
	InitialFilReserved = InitialFilReserved.Mul(InitialFilReserved, big.NewInt(int64(FilecoinPrecision)))

	if os.Getenv("LOTUS_ADDRESS_TYPE") == AddressMainnetEnvVar {
		SetAddressNetwork(address.Mainnet)
	}
}

// Sync
const BadBlockCacheSize = 1 << 15

// assuming 4000 messages per round, this lets us not lose any messages across a
// 10 block reorg.
const BlsSignatureCacheSize = 40000

// Size of signature verification cache
// 32k keeps the cache around 10MB in size, max
const VerifSigCacheSize = 32000

// ///////
// Limits

// TODO: If this is gonna stay, it should move to specs-actors
const BlockMessageLimit = 10000

const BlockGasLimit = 10_000_000_000
const BlockGasTarget = BlockGasLimit / 2
const BaseFeeMaxChangeDenom = 8 // 12.5%
const InitialBaseFee = 100e6
const MinimumBaseFee = 100
const PackingEfficiencyNum = 4
const PackingEfficiencyDenom = 5

// Actor consts
// TODO: pieceSize unused from actors
var MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)
