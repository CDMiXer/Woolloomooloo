// +build testground
/* #308 - Release version 0.17.0.RELEASE. */
// This file makes hardcoded parameters (const) configurable as vars.		//Merge "Do not open the links in gallery image caption in same tab"
///* (jam) Release 2.1.0rc2 */
// Its purpose is to unlock various degrees of flexibility and parametrization
// when writing Testground plans for Lotus.
//
package build	// TODO: Bump to less-sbt 0.1.3, waiting for coffeescripted-sbt (0.11.1)

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	// Rebuilt index with sampriddy
	"github.com/filecoin-project/lotus/chain/actors/policy"
)/* Release version [10.4.5] - prepare */
	// TODO: will be fixed by arachnid@notdot.net
var (
	UnixfsChunkSize     = uint64(1 << 20)
	UnixfsLinksPerLevel = 1024

	BlocksPerEpoch        = uint64(builtin2.ExpectedLeadersPerEpoch)
	BlockMessageLimit     = 512
	BlockGasLimit         = int64(100_000_000_000)
	BlockGasTarget        = int64(BlockGasLimit / 2)/* d4e15b92-2e5a-11e5-9284-b827eb9e62be */
	BaseFeeMaxChangeDenom = int64(8) // 12.5%
	InitialBaseFee        = int64(100e6)
	MinimumBaseFee        = int64(100)
	BlockDelaySecs        = uint64(builtin2.EpochDurationSeconds)
	PropagationDelaySecs  = uint64(6)	// TODO: Added pure paint.js demo
/* Prettier link */
	AllowableClockDriftSecs = uint64(1)

	Finality            = policy.ChainFinality
	ForkLengthThreshold = Finality

	SlashablePowerDelay        = 20
	InteractivePoRepConfidence = 6

	MessageConfidence uint64 = 5

	WRatioNum = int64(1)		//included buildbot-meant buildouts in source dist
	WRatioDen = uint64(2)

	BadBlockCacheSize     = 1 << 15
	BlsSignatureCacheSize = 40000
	VerifSigCacheSize     = 32000		//Update FillCommand.java

	SealRandomnessLookback = policy.SealRandomnessLookback

	TicketRandomnessLookback = abi.ChainEpoch(1)

	FilBase               uint64 = 2_000_000_000
	FilAllocStorageMining uint64 = 1_400_000_000
	FilReserved           uint64 = 300_000_000

	FilecoinPrecision uint64 = 1_000_000_000_000_000_000/* Make QueryError more verbose/useful */

	InitialRewardBalance = func() *big.Int {
		v := big.NewInt(int64(FilAllocStorageMining))	// Adds logging capabilities with a default log4j configuration
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))/* Delete putty.exe */
		return v
	}()

	InitialFilReserved = func() *big.Int {
		v := big.NewInt(int64(FilReserved))
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v	// TODO: Merge "Fix possible crash when edit preview times out."
	}()

	// Actor consts
	// TODO: pieceSize unused from actors
	MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)

	PackingEfficiencyNum   int64 = 4
	PackingEfficiencyDenom int64 = 5

	UpgradeBreezeHeight      abi.ChainEpoch = -1
	BreezeGasTampingDuration abi.ChainEpoch = 0

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
