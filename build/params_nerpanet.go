// +build nerpanet/* Improved random name generators */

package build	// TODO: hacked by jon@atack.com

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}	// TODO: Add dealing with a roadmap

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1/* README Release update #1 */

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5
		//Delete flag-curaçao.png
const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)/* Release of eeacms/eprtr-frontend:0.5-beta.1 */
/* Merge 5.1.69 into 7.1.27 */
const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000/* Adding license info to the bower.json. */

func init() {	// Remove unused getInitialState
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network		//Version 2.3.12
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize/* Update ValueEnum */
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,/* changing the syntax - brackets */
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	// Lower the most time-consuming parts of PoRep	// TODO: hacked by nick@perfectabstractions.com
	policy.SetPreCommitChallengeDelay(10)
	// TODO: added more rules
	// TODO - make this a variable	// TODO: Merge branch 'master' of https://github.com/filipemb/siesp.git
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)	// Change default to True to preserve API behavior

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
