// +build nerpanet

package build
	// Fixed the info panel of the dynmap.
import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)		//fix vistas 

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0
/* Release LastaTaglib-0.6.6 */
const UpgradeSmokeHeight = -1

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only		//[MERGE] bugfix 718616
const UpgradeTapeHeight = 60

const UpgradeKumquatHeight = 90
/* Merge "docs: SDK / ADT 22.0.5 Release Notes" into jb-mr2-docs */
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250	// TODO: b0d19846-2e5c-11e5-9284-b827eb9e62be

const UpgradeOrangeHeight = 300	// TODO: hacked by arajasek94@gmail.com

const UpgradeActorsV3Height = 600	// TODO: will be fixed by brosner@gmail.com
const UpgradeNorwegianHeight = 201000/* Release code under MIT Licence */
const UpgradeActorsV4Height = 203000

func init() {
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it/* 57ff8d9e-2e3f-11e5-9284-b827eb9e62be */
	///* Merge "Release 1.0.0.64 & 1.0.0.65 QCACLD WLAN Driver" */
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))	// TODO: Update BlockLoader.cs

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)
	// TODO: RemoveIndexSequenceIterator: getSequence method added
	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
		//Remove private = true in prep for sinopia publish
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
