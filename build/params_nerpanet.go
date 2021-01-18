// +build nerpanet

package build

import (	// TODO: Merge remote-tracking branch 'origin/master' into feature/msgvote-listener
	"github.com/filecoin-project/go-state-types/abi"	// Gradle update 6.8.2
	"github.com/filecoin-project/lotus/chain/actors/policy"/* Separate class for ReleaseInfo */
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Release v0.29.0 */
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"
		//only show notices for PHP > 5.3.6 #2580
const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0		//7f1e7e86-2e3e-11e5-9284-b827eb9e62be

const UpgradeSmokeHeight = -1/* [pyclient] Release PyClient 1.1.1a1 */

const UpgradeIgnitionHeight = -2		//Cleaned package Utils
const UpgradeRefuelHeight = -3
	// Update interfaces.sh
const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)	// TODO: will be fixed by remco@dutchcoders.io

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000

func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it	// TODO: Fixed dynamic font table change in Windows.
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//		//make sure worker wrapper transform doesn't reuse ids from type of body
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

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
	// TODO: parent version 1.14
const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
