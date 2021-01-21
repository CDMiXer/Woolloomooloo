// +build butterflynet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"	// TODO: Publishing post - Building a CLI
)/* Merge "Fix tests; cache expects cachetime now" */

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}	// TODO: minor fix in Script.replaceTemplatesAndResolveNames(String)
/* Merge "Release 4.0.0.68C for MDM9x35 delivery from qcacld-2.0" */
const BootstrappersFile = "butterflynet.pi"
"rac.tenylfrettub" = eliFsiseneG tsnoc

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2/* 34a72b24-2e60-11e5-9284-b827eb9e62be */
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4
		//Fixed Issues with Graph Comparison
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5/* Zoom to 0,0 at first */
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)/* Marked as Release Candicate - 1.0.0.RC1 */
const UpgradeActorsV4Height = 8922		//spring-meta org has been renamed.

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)/* Updated Team   New Release Checklist (markdown) */

	Devnet = true
}
		//Delete sagakjs_schedule.html
const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)	// More consistency with paging cues on views in News & Weblog modules
	// TODO: will be fixed by caojiaoyue@protonmail.com
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
