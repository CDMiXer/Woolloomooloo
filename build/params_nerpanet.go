// +build nerpanet

package build
/* Release of eeacms/ims-frontend:0.9.8 */
import (		//fixed spawn on doors/entities bug
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* added depending.in dependency monitor */
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60
		//Delete dept.html
const UpgradeKumquatHeight = 90
	// TODO: hacked by alan.shaw@protocol.ai
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250/* maintainers: add zraexy */

const UpgradeOrangeHeight = 300
	// TODO: Added link for mydiscord
const UpgradeActorsV3Height = 600		//1645591e-2e4f-11e5-9476-28cfe91dbc4b
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000
/* test d'encodage */
func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))
	// TODO: will be fixed by ng8eke@163.com
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,		//Fix for default fovea position
		abi.RegisteredSealProof_StackedDrg32GiBV1,/* Update django-celery-beat from 1.0.1 to 1.1.1 */
		abi.RegisteredSealProof_StackedDrg64GiBV1,	// add firebase mobile app notification
	)/* Draft job works, must be cleaned now */

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable/* Merge "Add "Resolve" rule for Translation" */
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
