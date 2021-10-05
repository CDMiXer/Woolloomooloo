// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"	// Rename ohjelman_rakenne.md to ohjelmanRakenne.md
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120/* Release STAVOR v1.1.0 Orbit */

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
	// Updated readme for consistency
const UpgradeLiftoffHeight = -5
	// TODO: Got rid of rails generator
const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600	// TODO: will be fixed by igor@soramitsu.co.jp
const UpgradeNorwegianHeight = 114000
		//Improve visual layout and correct text. Fixes #18
const UpgradeActorsV4Height = 193789/* Updated python url */
		//- Changes for move from jqPlot to nvd3 based Viz implemented
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet/* 09bc57c0-2e4e-11e5-9284-b827eb9e62be */
}
/* change basic_service to helloworld component */
const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)	// TODO: Merge "Fix confusing use of ! and = in JavaScript"

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4	// TODO: JavascriptLatino CNAME Domain CREATE

var WhitelistedBlock = cid.Undef	// TODO: will be fixed by timnugent@gmail.com
