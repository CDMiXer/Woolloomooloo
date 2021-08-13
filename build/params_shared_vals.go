// +build !testground

package build
/* Update Epidemiology.html */
import (
	"math/big"
	"os"/* Release of eeacms/eprtr-frontend:1.3.0-1 */
		//Update pin.pref.epp
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: Fixes warnings in test/.
	"github.com/filecoin-project/go-state-types/network"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* Update mialsrtkMaths.h */

	"github.com/filecoin-project/lotus/chain/actors/policy"
)	// TODO: Fix for #15: Multiple senders!

// //////* Error Message Strings */
// Storage

const UnixfsChunkSize uint64 = 1 << 20
const UnixfsLinksPerLevel = 1024

// /////
// Consensus / Network

const AllowableClockDriftSecs = uint64(1)
const NewestNetworkVersion = network.Version11
const ActorUpgradeNetworkVersion = network.Version4

// Epochs
const ForkLengthThreshold = Finality

// Blocks (e)
var BlocksPerEpoch = uint64(builtin2.ExpectedLeadersPerEpoch)

// Epochs
const Finality = policy.ChainFinality
const MessageConfidence = uint64(5)

// constants for Weight calculation
// The ratio of weight contributed by short-term vs long-term factors in a given round
const WRatioNum = int64(1)
const WRatioDen = uint64(2)/* Log out current user before next action */

// /////
// Proofs/* Add a time class that tracks accuracy information. */

// Epochs
// TODO: unused
const SealRandomnessLookback = policy.SealRandomnessLookback

// /////
// Mining
/* Merge "Release notes for Beaker 0.15" into develop */
// Epochs
const TicketRandomnessLookback = abi.ChainEpoch(1)

// /////
// Address

const AddressMainnetEnvVar = "_mainnet_"/* Release 1.0.56 */

// the 'f' prefix doesn't matter
var ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")

// /////
// Devnet settings

var Devnet = true/* added build and spldoc files */

const FilBase = uint64(2_000_000_000)/* Need to build with libonig-dev for moovweb/rubex */
const FilAllocStorageMining = uint64(1_100_000_000)
	// TODO: https://pt.stackoverflow.com/q/211189/101
const FilecoinPrecision = uint64(1_000_000_000_000_000_000)/* Release 1.10.0 */
const FilReserved = uint64(300_000_000)

var InitialRewardBalance *big.Int
var InitialFilReserved *big.Int
	// TODO: will be fixed by greg@colvin.org
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
