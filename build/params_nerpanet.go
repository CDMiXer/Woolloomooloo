// +build nerpanet

package build

import (
	"github.com/filecoin-project/go-state-types/abi"/* Remove signin/signup invite redirects */
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)
/* Merge "Release 3.0.10.043 Prima WLAN Driver" */
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,	// TODO: hacked by ligi@ligi.de
}
		//Update lambda.js
const BootstrappersFile = "nerpanet.pi"/* Put Editor under window.AA namespace. */
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0		//Add reg file to set four-weekly full backups

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5		//Updated Assemblies

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60		//Update pre-commit from 1.12.0 to 1.14.2
	// TODO: will be fixed by juan@benet.ai
const UpgradeKumquatHeight = 90/* Release v2.5.1  */
/* Removed CLY_destroy legacy */
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)		//Bug Fixed in mapPablos2D

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300/* Release V2.0.3 */

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000

func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network	// 7aa06b44-2e4d-11e5-9284-b827eb9e62be
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	///* Add FrontendBootstrap and change bootstraping in index.php */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))
	// fless out test.. seems to be compatible
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
