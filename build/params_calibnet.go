// +build calibnet

package build
	// TODO: Merge "Add NetworkAndCompute Lister and ShowOne classes"
import (
	"github.com/filecoin-project/go-address"/* Release 13.0.0 */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"/* Fix Zon wizard sip id */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}	// TODO: will be fixed by witek@enjin.io

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
	// TODO: will be fixed by lexy8russo@outlook.com
const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3		//LineWrap example explains limitations when using 16x1 / 8x2 displays
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
/* Released 1.6.0-RC1. */
const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90
/* 854e48fa-2e47-11e5-9284-b827eb9e62be */
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)/* Release of eeacms/forests-frontend:1.8.2 */

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300
		//Change final styling and some of the main.html layout.
const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000
/* Merge "Release 1.0.0.146 QCACLD WLAN Driver" */
const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}/* add MiniMusicCmdLine */
/* Tagged Release 2.1 */
const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)/* fixed word coords matching for hyphenated words; refs #18536 */

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start/* Merged branch develop into feature/#23 */
const BootstrapPeerThreshold = 4
	// Update 2. Commands.md
var WhitelistedBlock = cid.Undef
