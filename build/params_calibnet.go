// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)	// TODO: hacked by hello@brooklynzelenka.com
		//Added y axis.
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120	// TODO: first pass at AJAX

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3	// Tweak docs per #73
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
	// TODO: Added marquee selection to scene editor.
const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100		//Added deps to pod spec
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)	// Fixing analytics code.

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000/* Release v0.1.3 */

const UpgradeActorsV4Height = 193789
/* Release 1.0.29 */
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))/* Fix pydev project: remove hardcoded reference to the requests library  */
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,/* Release 8.0.2 */
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
		//Part 2 of recreating license
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start/* Utilisation de la classe entrepot dans planche de jeu avec la méthode recolte */
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
