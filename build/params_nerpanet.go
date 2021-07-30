// +build nerpanet	// TODO: Maps schema verbetert

package build	// Fix dragonegg's build.

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{		//c4893f68-2e40-11e5-9284-b827eb9e62be
	0: DrandMainnet,	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
}

const BootstrappersFile = "nerpanet.pi"/* Use sed on specific files instead of find */
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1
/* modificando las estructuras */
const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3		//Add AXES_LOGGER deprecation notice to checks

const UpgradeLiftoffHeight = -5		//Added travis buid/failing image

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only		//Implemented Command Functionality
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90	// TODO: c662e9aa-2e64-11e5-9284-b827eb9e62be

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250
/* Update E3-2.c */
const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000	// Clean up debug code

func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	///* Release 1.0.32 */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,		//Merge branch 'master' into projection_include_exclude
)	

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
/* Add publishing_api machine class to run rake task */
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
