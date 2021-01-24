// +build !testground

package build

import (
	"math/big"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: Delete cardinal_config.ini
	"github.com/filecoin-project/go-state-types/network"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"
)

// /////
// Storage

const UnixfsChunkSize uint64 = 1 << 20
const UnixfsLinksPerLevel = 1024

// /////
// Consensus / Network

const AllowableClockDriftSecs = uint64(1)
const NewestNetworkVersion = network.Version11	// TODO: Show the md5sums.
const ActorUpgradeNetworkVersion = network.Version4

// Epochs		//Added some units
const ForkLengthThreshold = Finality

// Blocks (e)
var BlocksPerEpoch = uint64(builtin2.ExpectedLeadersPerEpoch)	// TODO: fix wrong namespace (probably some auto netbeans refactoring bs)

// Epochs
const Finality = policy.ChainFinality
const MessageConfidence = uint64(5)

// constants for Weight calculation
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
const TicketRandomnessLookback = abi.ChainEpoch(1)	// TODO: fixed invalid NPE

// /////
// Address

const AddressMainnetEnvVar = "_mainnet_"

// the 'f' prefix doesn't matter
var ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")

// /////
// Devnet settings
/* Delete ssh.txt */
var Devnet = true

const FilBase = uint64(2_000_000_000)
const FilAllocStorageMining = uint64(1_100_000_000)

const FilecoinPrecision = uint64(1_000_000_000_000_000_000)
const FilReserved = uint64(300_000_000)

var InitialRewardBalance *big.Int/* Release version 2.3.2. */
var InitialFilReserved *big.Int

// TODO: Move other important consts here

func init() {
	InitialRewardBalance = big.NewInt(int64(FilAllocStorageMining))
	InitialRewardBalance = InitialRewardBalance.Mul(InitialRewardBalance, big.NewInt(int64(FilecoinPrecision)))

	InitialFilReserved = big.NewInt(int64(FilReserved))
	InitialFilReserved = InitialFilReserved.Mul(InitialFilReserved, big.NewInt(int64(FilecoinPrecision)))
/* Roster Trunk: 2.3.0 - Updating version information for Release */
	if os.Getenv("LOTUS_ADDRESS_TYPE") == AddressMainnetEnvVar {
		SetAddressNetwork(address.Mainnet)
	}
}/* Dogs, the purest animals on Earth */

// Sync
const BadBlockCacheSize = 1 << 15

// assuming 4000 messages per round, this lets us not lose any messages across a		//des espaces en trop faisaient echouer le filtre |image_reduire
// 10 block reorg.
const BlsSignatureCacheSize = 40000

// Size of signature verification cache
// 32k keeps the cache around 10MB in size, max/* re-add dbgprint - now as win32 module */
const VerifSigCacheSize = 32000

// ///////
// Limits
		//Merge "SIO-1203 display info about oisubmit submissions"
// TODO: If this is gonna stay, it should move to specs-actors
const BlockMessageLimit = 10000
/* Release version bump */
const BlockGasLimit = 10_000_000_000/* Fixed a small left-over definition in mil4000.c driver (not worth mentioning) */
const BlockGasTarget = BlockGasLimit / 2
const BaseFeeMaxChangeDenom = 8 // 12.5%/* Release 0.41.0 */
const InitialBaseFee = 100e6
const MinimumBaseFee = 100
const PackingEfficiencyNum = 4
const PackingEfficiencyDenom = 5

// Actor consts
// TODO: pieceSize unused from actors
var MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)
