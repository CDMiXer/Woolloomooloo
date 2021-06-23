// +build nerpanet

package build

import (
	"github.com/filecoin-project/go-state-types/abi"/* Option for bin2vmem to load the bin file at an address != 0x400 */
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"		//added helper to find all methods

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"/* Smaller fix */

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1
/* Added python-pil to the list of prerequsites */
const UpgradeIgnitionHeight = -2/* Added EncodeDecodeTest.java */
const UpgradeRefuelHeight = -3
		//ADD symfony2 framework 2.4.1 -- basic version
const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only/* added basic library creation support */
const UpgradeTapeHeight = 60
/* Merge "Release 3.2.3.443 Prima WLAN Driver" */
const UpgradeKumquatHeight = 90		//TODO-632: ditching template fun for now

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000

func init() {/* [Release] 5.6.3 */
	// Minimum block production power is set to 4 TiB
	// Rationale is to discourage small-scale miners from trying to take over the network		//f6a99b22-2e69-11e5-9284-b827eb9e62be
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//	// TODO: Merge branch 'master' into spritetext-precache
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize	// refresh cache and force update every 60 mins to hasten pickup of updates
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)	// TODO: will be fixed by davidad@alum.mit.edu

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable/* Submit mquiz results to server */
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false	// TODO: will be fixed by greg@colvin.org
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
