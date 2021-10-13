// +build nerpanet

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"		//Create compare2lists.py
)		//Integrated origin/master

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}
/* Update MacCoder.java */
const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1	// TODO: will be fixed by souzau@yandex.com
const BreezeGasTampingDuration = 0
/* Define that we'll use the MIT license. */
const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only/* TASK: Add installation hint to README.md */
const UpgradeTapeHeight = 60/* Add -p parameter to create parent folders. */
/* Remove extra } */
const UpgradeKumquatHeight = 90
		//Allow using non-table alias as a rhs relation name, fix for #84 and #59
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000/* Tagging a Release Candidate - v3.0.0-rc17. */

func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize/* rollback of failed algo for exploring. */
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))/* Release 4.7.3 */

	policy.SetSupportedProofTypes(		//Delete configure-chroot~
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)		//Merge "Allow overriding pacemaker_node_ips for staged upgrade"

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)
/* Fix tests now that "lifetimes" is commented out */
const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
