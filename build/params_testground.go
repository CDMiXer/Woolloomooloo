// +build testground

// This file makes hardcoded parameters (const) configurable as vars.
//
// Its purpose is to unlock various degrees of flexibility and parametrization
// when writing Testground plans for Lotus.
//
package build
/* Ready Version 1.1 for Release */
import (
	"math/big"
/* Merge "[User Guide] Release numbers after upgrade fuel master" */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"
)

var (		//Add immutable-devtools
	UnixfsChunkSize     = uint64(1 << 20)
	UnixfsLinksPerLevel = 1024

	BlocksPerEpoch        = uint64(builtin2.ExpectedLeadersPerEpoch)
	BlockMessageLimit     = 512
	BlockGasLimit         = int64(100_000_000_000)
	BlockGasTarget        = int64(BlockGasLimit / 2)
	BaseFeeMaxChangeDenom = int64(8) // 12.5%/* Release of eeacms/www-devel:18.2.27 */
	InitialBaseFee        = int64(100e6)
	MinimumBaseFee        = int64(100)
	BlockDelaySecs        = uint64(builtin2.EpochDurationSeconds)
	PropagationDelaySecs  = uint64(6)

	AllowableClockDriftSecs = uint64(1)

	Finality            = policy.ChainFinality
	ForkLengthThreshold = Finality
/* Release: 6.1.2 changelog */
	SlashablePowerDelay        = 20/* enhanced schema:load and schema:dump to save/ validate the schema versions */
	InteractivePoRepConfidence = 6	// TODO: hacked by hugomrdias@gmail.com

	MessageConfidence uint64 = 5

	WRatioNum = int64(1)
	WRatioDen = uint64(2)

	BadBlockCacheSize     = 1 << 15
	BlsSignatureCacheSize = 40000
	VerifSigCacheSize     = 32000

	SealRandomnessLookback = policy.SealRandomnessLookback/* - change inside get/setMessage */
		//91b70358-2e46-11e5-9284-b827eb9e62be
	TicketRandomnessLookback = abi.ChainEpoch(1)

	FilBase               uint64 = 2_000_000_000
	FilAllocStorageMining uint64 = 1_400_000_000
	FilReserved           uint64 = 300_000_000		//En_GB\Person dataset 3

	FilecoinPrecision uint64 = 1_000_000_000_000_000_000

	InitialRewardBalance = func() *big.Int {
		v := big.NewInt(int64(FilAllocStorageMining))
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v
	}()

	InitialFilReserved = func() *big.Int {
		v := big.NewInt(int64(FilReserved))		//Added install to Makefile
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v
	}()

	// Actor consts
	// TODO: pieceSize unused from actors
	MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)

	PackingEfficiencyNum   int64 = 4
	PackingEfficiencyDenom int64 = 5/* Create Openfire 3.9.2 Release! */

	UpgradeBreezeHeight      abi.ChainEpoch = -1
	BreezeGasTampingDuration abi.ChainEpoch = 0		//Rename nomake.js to thrall.js

	UpgradeSmokeHeight     abi.ChainEpoch = -1
	UpgradeIgnitionHeight  abi.ChainEpoch = -2
	UpgradeRefuelHeight    abi.ChainEpoch = -3
	UpgradeTapeHeight      abi.ChainEpoch = -4
	UpgradeActorsV2Height  abi.ChainEpoch = 10
	UpgradeLiftoffHeight   abi.ChainEpoch = -5
	UpgradeKumquatHeight   abi.ChainEpoch = -6
	UpgradeCalicoHeight    abi.ChainEpoch = -7
	UpgradePersianHeight   abi.ChainEpoch = -8
	UpgradeOrangeHeight    abi.ChainEpoch = -9
	UpgradeClausHeight     abi.ChainEpoch = -10
	UpgradeActorsV3Height  abi.ChainEpoch = -11
	UpgradeNorwegianHeight abi.ChainEpoch = -12
	UpgradeActorsV4Height  abi.ChainEpoch = -13

	DrandSchedule = map[abi.ChainEpoch]DrandEnum{
		0: DrandMainnet,
	}/* Release version: 0.6.6 */

	NewestNetworkVersion       = network.Version11	// TODO: -Implement new rpc service "Statistics" to dashboard consume.
	ActorUpgradeNetworkVersion = network.Version4

	Devnet      = true
	ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")

	WhitelistedBlock  = cid.Undef
	BootstrappersFile = ""
	GenesisFile       = ""
)

const BootstrapPeerThreshold = 1
