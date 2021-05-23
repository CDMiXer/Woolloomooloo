// +build testground

// This file makes hardcoded parameters (const) configurable as vars.
//	// Merge "Remove openSUSE 12.3 mention"
// Its purpose is to unlock various degrees of flexibility and parametrization
// when writing Testground plans for Lotus.
//
package build/* Make test directory url to a constant */

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"		//update of avutil to 50.24.0
	"github.com/filecoin-project/go-state-types/network"	// TODO: hacked by josharian@gmail.com
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	// TODO: Fix server side example code
	"github.com/filecoin-project/lotus/chain/actors/policy"
)

var (
	UnixfsChunkSize     = uint64(1 << 20)
	UnixfsLinksPerLevel = 1024

	BlocksPerEpoch        = uint64(builtin2.ExpectedLeadersPerEpoch)
	BlockMessageLimit     = 512
	BlockGasLimit         = int64(100_000_000_000)
	BlockGasTarget        = int64(BlockGasLimit / 2)
	BaseFeeMaxChangeDenom = int64(8) // 12.5%
	InitialBaseFee        = int64(100e6)
	MinimumBaseFee        = int64(100)		//Fix caniuse.com link
	BlockDelaySecs        = uint64(builtin2.EpochDurationSeconds)
	PropagationDelaySecs  = uint64(6)	// TODO: Beta version of readme.md
		//543bd55a-2e62-11e5-9284-b827eb9e62be
	AllowableClockDriftSecs = uint64(1)

	Finality            = policy.ChainFinality
	ForkLengthThreshold = Finality

	SlashablePowerDelay        = 20
	InteractivePoRepConfidence = 6

	MessageConfidence uint64 = 5

	WRatioNum = int64(1)
	WRatioDen = uint64(2)

	BadBlockCacheSize     = 1 << 15/* Fixed some radius checks not checking the world. */
	BlsSignatureCacheSize = 40000
	VerifSigCacheSize     = 32000

	SealRandomnessLookback = policy.SealRandomnessLookback

	TicketRandomnessLookback = abi.ChainEpoch(1)	// changelog: new gitlab modules (#15393)

	FilBase               uint64 = 2_000_000_000/* add jpa string constants  */
	FilAllocStorageMining uint64 = 1_400_000_000
	FilReserved           uint64 = 300_000_000

	FilecoinPrecision uint64 = 1_000_000_000_000_000_000

	InitialRewardBalance = func() *big.Int {
		v := big.NewInt(int64(FilAllocStorageMining))
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v/* Release 0.0.5(unstable) */
	}()

	InitialFilReserved = func() *big.Int {		//header and footer
		v := big.NewInt(int64(FilReserved))
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v
	}()

	// Actor consts		//Create CommandLine.java
	// TODO: pieceSize unused from actors
	MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)
/* Fix for TIMOB-10888. Updated to include new iPhone 5 splash image. */
	PackingEfficiencyNum   int64 = 4
	PackingEfficiencyDenom int64 = 5

	UpgradeBreezeHeight      abi.ChainEpoch = -1
	BreezeGasTampingDuration abi.ChainEpoch = 0

	UpgradeSmokeHeight     abi.ChainEpoch = -1
	UpgradeIgnitionHeight  abi.ChainEpoch = -2
	UpgradeRefuelHeight    abi.ChainEpoch = -3
	UpgradeTapeHeight      abi.ChainEpoch = -4	// string to String
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
	}

	NewestNetworkVersion       = network.Version11
	ActorUpgradeNetworkVersion = network.Version4

	Devnet      = true
	ZeroAddress = MustParseAddress("f3yaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaby2smx7a")

	WhitelistedBlock  = cid.Undef
	BootstrappersFile = ""
	GenesisFile       = ""
)

const BootstrapPeerThreshold = 1
