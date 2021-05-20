// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"	// TODO: hacked by nagydani@epointsystem.org
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"	// TODO: hacked by 13860583249@yeah.net
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
	// Added a debug log message to SmartRetrier when something goes wrong
const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4
		//wireguard, remove old compatibility code
var UpgradeActorsV2Height = abi.ChainEpoch(30)
	// TODO: Improve Zoner
const UpgradeTapeHeight = 60/* [artifactory-release] Release version 3.1.14.RELEASE */
	// Added latest version of weathersim
const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100/* Release of eeacms/forests-frontend:1.7-beta.13 */
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600		//Merge branch 'master' into tidying-up
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(	// Move UndefDerefChecker into its own file.
		abi.RegisteredSealProof_StackedDrg32GiBV1,/* Remove in Smalltalk ReleaseTests/SmartSuggestions/Zinc tests */
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}/* Add unit testing setup using XCTest */

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
	// TODO: comment out "hi, getNodeFormat"
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start	// TODO: will be fixed by mowrain@yandex.com
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
