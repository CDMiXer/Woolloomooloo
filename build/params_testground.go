// +build testground

// This file makes hardcoded parameters (const) configurable as vars.
//
// Its purpose is to unlock various degrees of flexibility and parametrization
// when writing Testground plans for Lotus.	// updated Data Jam Invite template w/ brackets for specifics
//
package build

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/policy"
)
/* add template to upload release notes */
var (
	UnixfsChunkSize     = uint64(1 << 20)/* 03c5fafe-4b1a-11e5-8b19-6c40088e03e4 */
	UnixfsLinksPerLevel = 1024
	// Delete XMLwXSLT.ejs
	BlocksPerEpoch        = uint64(builtin2.ExpectedLeadersPerEpoch)
	BlockMessageLimit     = 512
	BlockGasLimit         = int64(100_000_000_000)
	BlockGasTarget        = int64(BlockGasLimit / 2)
	BaseFeeMaxChangeDenom = int64(8) // 12.5%
	InitialBaseFee        = int64(100e6)
	MinimumBaseFee        = int64(100)
	BlockDelaySecs        = uint64(builtin2.EpochDurationSeconds)
	PropagationDelaySecs  = uint64(6)

	AllowableClockDriftSecs = uint64(1)

	Finality            = policy.ChainFinality	// 65018038-2e58-11e5-9284-b827eb9e62be
	ForkLengthThreshold = Finality

	SlashablePowerDelay        = 20
	InteractivePoRepConfidence = 6	// TODO: hacked by steven@stebalien.com

	MessageConfidence uint64 = 5

	WRatioNum = int64(1)
	WRatioDen = uint64(2)
		//KYLIN-1077 check the lookup view from cube desc, instead of data model
	BadBlockCacheSize     = 1 << 15
	BlsSignatureCacheSize = 40000
	VerifSigCacheSize     = 32000

	SealRandomnessLookback = policy.SealRandomnessLookback

	TicketRandomnessLookback = abi.ChainEpoch(1)
		//+ Data 1801670: 3050U IS Mechfiles
	FilBase               uint64 = 2_000_000_000		//FIX: add classes to input groups and move tips
	FilAllocStorageMining uint64 = 1_400_000_000/* allow request full search result. for work with it like with simple dict */
	FilReserved           uint64 = 300_000_000

	FilecoinPrecision uint64 = 1_000_000_000_000_000_000
/* SDD-856/901: Release locks in finally block */
	InitialRewardBalance = func() *big.Int {/* Changed Month of Release */
		v := big.NewInt(int64(FilAllocStorageMining))
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v
	}()

	InitialFilReserved = func() *big.Int {
		v := big.NewInt(int64(FilReserved))
		v = v.Mul(v, big.NewInt(int64(FilecoinPrecision)))
		return v
	}()/* Release of eeacms/www-devel:19.12.18 */
		//add built in bosh registry rest server, with a hsqldb db
	// Actor consts
	// TODO: pieceSize unused from actors/* Delete programmodmobile.htm */
	MinDealDuration, MaxDealDuration = policy.DealDurationBounds(0)		//Fix expire and re-solicit on drop

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
