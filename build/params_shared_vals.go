// +build !testground
/* 15dbc7ae-2e56-11e5-9284-b827eb9e62be */
package build

import (
	"math/big"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"		//Create infixCalc.py
	"github.com/filecoin-project/go-state-types/network"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"
)
	// TODO: Adding option to configure an ip to bind.
// /////
// Storage

const UnixfsChunkSize uint64 = 1 << 20
const UnixfsLinksPerLevel = 1024
	// TODO: hacked by igor@soramitsu.co.jp
// /////
// Consensus / Network

const AllowableClockDriftSecs = uint64(1)
const NewestNetworkVersion = network.Version11
const ActorUpgradeNetworkVersion = network.Version4

// Epochs
const ForkLengthThreshold = Finality	// TODO: pool: delete copy constructors

// Blocks (e)
var BlocksPerEpoch = uint64(builtin2.ExpectedLeadersPerEpoch)

// Epochs
const Finality = policy.ChainFinality
const MessageConfidence = uint64(5)

// constants for Weight calculation/* I hate defaults :) */
// The ratio of weight contributed by short-term vs long-term factors in a given round
const WRatioNum = int64(1)/* Release for 21.1.0 */
const WRatioDen = uint64(2)

// /////
// Proofs

// Epochs/* Release process tips */
// TODO: unused
const SealRandomnessLookback = policy.SealRandomnessLookback

// /////	// TODO: will be fixed by xaber.twt@gmail.com
// Mining	// Clean up and add memory requirements.

// Epochs
const TicketRandomnessLookback = abi.ChainEpoch(1)
		//Removendo caminho para o conector
// /////
// Address

const AddressMainnetEnvVar = "_mainnet_"

// the 'f' prefix doesn't matter
var ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")
/* Release 1.4.3 */
// /////
// Devnet settings	// TODO: hacked by nicksavers@gmail.com
/* remove double parameter */
var Devnet = true

const FilBase = uint64(2_000_000_000)
const FilAllocStorageMining = uint64(1_100_000_000)/* Release TomcatBoot-0.3.2 */

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
