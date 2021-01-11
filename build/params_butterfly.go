// +build butterflynet

package build

import (/* Gradle Release Plugin - new version commit:  "2.7-SNAPSHOT". */
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{	// TODO: update package information.
	0: DrandMainnet,
}/* Merge "[INTERNAL][FIX] unified.FileUploader height in toolbar fixed" */

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120		//Build 45a: Add Junit Tests, Removed Client/Server code.
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4/* GROOVY-2445: Provide an enhanced each() method for subclasses of java.lang.Enum */
		//Delete Point.h.gch
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60		//pt-osc: Change --quiet back and remove the quietness checks
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120/* removing unused indexes */
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)	// Fix NOTICEs files wrt jline
const UpgradeActorsV4Height = 8922

func init() {		//wall collision 
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(/* Delete Acuaticas.java */
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true
}	// Merge "[INTERNAL][FIX] uxap.ObjectPage AnchorBar HCW, HCB styling fixed"

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)
/* Merge "Release 3.2.3.414 Prima WLAN Driver" */
const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start		//Atualiza dados da coleção West Indies
const BootstrapPeerThreshold = 2/* 686405a0-2e5a-11e5-9284-b827eb9e62be */

var WhitelistedBlock = cid.Undef
