// +build calibnet
	// TODO: Update home-view.json
package build

import (/* Release tag: 0.6.4. */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"/* bec2391a-2e48-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)/* Release of Milestone 3 of 1.7.0 */

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{	// TODO: hacked by why@ipfs.io
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"
/* Release fixed. */
const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2/* actual content, update on version 1.0.1 */
		//xfail for matplotlib bug with SVG output
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)		//Add new Lucee5 functions and tags to listing
		//e4c5e28a-2e6e-11e5-9284-b827eb9e62be
const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5		//MetricSchemasF: drop event if size > 64000
/* add more javadoc. */
const UpgradeKumquatHeight = 90/* Renamed member variable to make code clearer. */

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250/* [Release] 5.6.3 */

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789
	// Merge branch 'master' into dev_27
func init() {/* remove out of date "where work is happening" and link to Releases page */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
