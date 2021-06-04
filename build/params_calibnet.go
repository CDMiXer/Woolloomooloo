// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"		//Merge "Implements field validation for complex query functionality"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)
		//Уточнение вопроса с HTML5.
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"		//Ignored intermitted test failure
const GenesisFile = "calibnet.car"	// TODO: Add possibility to change final filename

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2		//Catch SF BUG 1621938: gimpact only does stride 12.

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
/* Release for 18.28.0 */
const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5
/* Updated PiAware Release Notes (markdown) */
const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

052 = thgieHsualCedargpU tsnoc

const UpgradeOrangeHeight = 300/* Release Roadmap */

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000
/* Release of eeacms/www:18.3.23 */
const UpgradeActorsV4Height = 193789

func init() {	// ispravka fill funkcije
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet/* add doc for ff_osc ugens */
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)/* Moved the javascript to the js directory. */

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
