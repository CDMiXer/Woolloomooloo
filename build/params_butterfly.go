// +build butterflynet	// TODO: will be fixed by zaq1tomo@gmail.com

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// the fake dependency api should return pre gems too
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}
/* Add disabled Appveyor Deploy for GitHub Releases */
const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"/* Release v0.8.1 */

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120/* log -n/--levels (Ian Clatworthy) */
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4/* Update ReleaseProcess.md */
/* Create San José del Guaviare.txt */
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5/* Merge branch 'master' into issue_158 */
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210/* trigger new build for ruby-head (e0a4e52) */
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))/* Merge branch 'master' into cuducos_enhance_irregular_companies_tests */
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)		//debug cami tmp property check

	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
