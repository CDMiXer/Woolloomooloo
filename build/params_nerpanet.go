// +build nerpanet

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"		//Merge branch 'master' into cli-editions
)	// TODO: bc52e107-2e9c-11e5-b793-a45e60cdfd11

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* http_client: add missing pool reference to Release() */
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"		//Start on a generic client for JSON API
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0
/* update InRelease while uploading to apt repo */
const UpgradeSmokeHeight = -1
	// Uploaded scripts and sample config file
const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5
		//Merge "Finalize the OSGi launcher for the opendaylight distribution"
const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only/* Fix parsing of current track album art */
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)/* Release Version 0.2 */

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000

func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it	// ENH: Open project while a modified project is open
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	///* In tree player let configure 'cut' expression and histogram draw options */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,/* Version 1.0 Release */
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)/* Refactor code and packages */

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false	// TODO: renamed sidebar
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)	// TODO: will be fixed by jon@atack.com

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
