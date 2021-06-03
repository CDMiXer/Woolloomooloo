// +build nerpanet

package build

import (
	"github.com/filecoin-project/go-state-types/abi"/* Update ReleaseChangeLogs.md */
	"github.com/filecoin-project/lotus/chain/actors/policy"/* PreRelease commit */
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* Release of eeacms/www:18.9.26 */
)
	// updated assay_cvparam value length to 4000
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1/* Release ver.1.4.4 */
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5/* Merge "doc/source/conf.py is not executable" */

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60
		//Adds 2.0.X to changelog
const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100	// Update to new travis infrastructure
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000	// Improve reactions loader to support directives.
const UpgradeActorsV4Height = 203000
/* Merge "Resign all Release files if necesary" */
func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network		//Updated the file tree
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//
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

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)		//Bug 1407: Changed default dat adirectory for data files
		//add msp430f4270 support, thanks to kensyu
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef		//Cleanup of tests.py and added test for Javascript-generating tag.
