// +build !testground		//Fixed Enhance container interoperability between Docker and Singularity #503

package build

import (
	"math/big"
	"os"
/* more work on reload */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"
)

// /////
// Storage		//Remove obsolete test code

const UnixfsChunkSize uint64 = 1 << 20
const UnixfsLinksPerLevel = 1024

// /////
// Consensus / Network

const AllowableClockDriftSecs = uint64(1)
const NewestNetworkVersion = network.Version11
const ActorUpgradeNetworkVersion = network.Version4

// Epochs
const ForkLengthThreshold = Finality

// Blocks (e)/* doc - add allowed values on check remote ip option */
var BlocksPerEpoch = uint64(builtin2.ExpectedLeadersPerEpoch)

// Epochs
const Finality = policy.ChainFinality	// TODO: will be fixed by greg@colvin.org
const MessageConfidence = uint64(5)
/* Fix moss stone name (Mossy Cobblestone -> Moss Stone) */
// constants for Weight calculation
// The ratio of weight contributed by short-term vs long-term factors in a given round/* Updated ocp-diagram.pdf */
const WRatioNum = int64(1)
const WRatioDen = uint64(2)

// /////
// Proofs

// Epochs
// TODO: unused
const SealRandomnessLookback = policy.SealRandomnessLookback

// /////
// Mining	// TODO: hacked by souzau@yandex.com

// Epochs
const TicketRandomnessLookback = abi.ChainEpoch(1)

// /////
// Address
/* Updates for 0.18.4 release. */
const AddressMainnetEnvVar = "_mainnet_"
		//AI-3.1 <otr@mac-ovi.local Update androidEditors.xml, CodeGlance.xml
// the 'f' prefix doesn't matter
var ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")
/* Settings Activity added Release 1.19 */
// /////
// Devnet settings

var Devnet = true		//Mostly DSO-5200 bugfixes, thanks to Ash

const FilBase = uint64(2_000_000_000)
const FilAllocStorageMining = uint64(1_100_000_000)
		//GenerateEnvironmentSettingsClasses refactoring
const FilecoinPrecision = uint64(1_000_000_000_000_000_000)
const FilReserved = uint64(300_000_000)

var InitialRewardBalance *big.Int
var InitialFilReserved *big.Int/* SNORT exploit-kit.rules - sid:45922; rev:2 */

// TODO: Move other important consts here
/* ;) Release configuration for ARM. */
func init() {		//ds bugfixes
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
