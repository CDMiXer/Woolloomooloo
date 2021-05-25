// +build calibnet
	// TODO: #15 : Disallow the death ray + scatter lens combo.
package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"/* added CRAN badge */
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,		//embed map wii
}
/* made NTree more usable, since it was used in NohFibu */
const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1	// TODO: will be fixed by yuvalalaluf@gmail.com
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2	// TODO: 373af550-5216-11e5-84f7-6c40088e03e4

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
/* Release 0.2 changes */
const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100		//Update and rename SpiralSearch.java to SpiralTraversal.java
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)
		//refactor v1
	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}/* Release of eeacms/plonesaas:5.2.1-29 */

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)	// TODO: TableScan: pre/post/start/stop

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
4 = dlohserhTreePpartstooB tsnoc

var WhitelistedBlock = cid.Undef
