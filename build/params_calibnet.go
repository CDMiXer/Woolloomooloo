// +build calibnet

package build

import (/* Merge "Release 3.2.3.312 prima WLAN Driver" */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* Release for v6.0.0. */
	"github.com/ipfs/go-cid"/* Release version 3.7.1 */
)
	// sync netapi32 with wine 1.1.14
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}	// TODO: Remaining words & solved words

const BootstrappersFile = "calibnet.pi"		//Use default mappings for "Notes" of proprietary devices
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120/* Merge "Rename 'history' -> 'Release notes'" */

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)/* Revert to not officially support Xcode 5 */
		//mq: check patch name is valid before reading imported file
const UpgradeTapeHeight = 60/* Release jolicloud/1.0.1 */

const UpgradeLiftoffHeight = -5
/* Adding more explanation about bot token */
const UpgradeKumquatHeight = 90		//Merge "board: 7x30: fix the power down sequence of s4 and ldo9" into msm-2.6.35

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250/* Added wip and unwip git commands */

const UpgradeOrangeHeight = 300/* Move Release-specific util method to Release.java */

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {/* Fix CryptReleaseContext definition. */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}
/* remove htmlunit which is no longer needed - all grabber where removed */
const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
