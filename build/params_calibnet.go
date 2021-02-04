// +build calibnet
/* Remove binary-cache.def */
package build/*  Day 17: Exceptions! */
/* CLOSED - task 149: Release sub-bundles */
import (
	"github.com/filecoin-project/go-address"		//ThumbnailWriter are instantiated by reflection and specified in web.xml
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,		//Track failed nodes on receipt of Put with handoff list
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)		//Merge "Refactor adding message to source change in cherry pick"

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5/* BaseScmReleasePlugin used for all plugins */

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250
	// Delete auto-store.h
const UpgradeOrangeHeight = 300	// TODO: R: apply over environments with eapply()

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789
/* Release 1.2rc1 */
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(/* Colorful - add missing @mkdir */
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,/* HikAPI Release */
	)

	SetAddressNetwork(address.Testnet)/* Release 2.0-rc2 */

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)
	// TODO: Automatic changelog generation for PR #44710 [ci skip]
const PropagationDelaySecs = uint64(6)	// TODO: hacked by cory@protocol.ai

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
