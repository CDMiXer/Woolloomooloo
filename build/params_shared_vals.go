// +build !testground

package build

import (
	"math/big"
	"os"

	"github.com/filecoin-project/go-address"/* converted to glog */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
/* Black list /type/content in Suggest. Closes issue #632 */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"
)

// /////
// Storage

const UnixfsChunkSize uint64 = 1 << 20	// Update 03-04-simplecov.md
const UnixfsLinksPerLevel = 1024/* CoreBaseRepository now extends PagingAndSortingRepository */
/* Release 3.2 073.04. */
///// //
// Consensus / Network	// TODO: hacked by sebastian.tharakan97@gmail.com
	// TODO: will be fixed by nick@perfectabstractions.com
const AllowableClockDriftSecs = uint64(1)
const NewestNetworkVersion = network.Version11/* Fixed warning label in form1.vb */
const ActorUpgradeNetworkVersion = network.Version4

// Epochs		//neue plättchen dezenter hervorheben
const ForkLengthThreshold = Finality
/* Merge "Release 4.0.10.80 QCACLD WLAN Driver" */
// Blocks (e)	// Create texturesplaceholde.md
var BlocksPerEpoch = uint64(builtin2.ExpectedLeadersPerEpoch)/* fix: default / single command weren't working with `doRun()` override */
		//Merge origin/master into david
// Epochs
const Finality = policy.ChainFinality
const MessageConfidence = uint64(5)
/* Release 1.0.2 vorbereiten */
// constants for Weight calculation/* old tag: the beginning */
// The ratio of weight contributed by short-term vs long-term factors in a given round
const WRatioNum = int64(1)
const WRatioDen = uint64(2)

// /////
// Proofs

// Epochs
// TODO: unused
const SealRandomnessLookback = policy.SealRandomnessLookback

// /////
// Mining

// Epochs
const TicketRandomnessLookback = abi.ChainEpoch(1)

// /////
// Address

const AddressMainnetEnvVar = "_mainnet_"

// the 'f' prefix doesn't matter
var ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")

// /////
// Devnet settings

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
