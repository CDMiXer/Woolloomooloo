// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"/* Updated documentation and website. Release 1.1.1. */
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,	// TODO: 06739088-2e49-11e5-9284-b827eb9e62be
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"		//Try even further measures in getting it to work

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120		//Some progress on issue 52... - issue 52

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4	// TODO: Aplicación SmartThing Web

var UpgradeActorsV2Height = abi.ChainEpoch(30)/* bugfix release 2.2.1 and prepare new release 2.2.2 */
/* Triggering also Busy Emotion. (Possible OpenNARS-1.6.3 Release Commit?) */
const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300
		//8c4d246a-2e41-11e5-9284-b827eb9e62be
const UpgradeActorsV3Height = 600/* Release of eeacms/jenkins-slave:3.12 */
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))/* Fix modified_since */
	policy.SetSupportedProofTypes(/* Released roombooking-1.0.0.FINAL */
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)	// added ability to modify instance write interval for testing

	Devnet = true

	BuildType = BuildCalibnet
}
/* Merge "[FIX] Parse/Validation messages not added/cleared correcly." */
const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4
/* Update misc dev deps */
var WhitelistedBlock = cid.Undef
