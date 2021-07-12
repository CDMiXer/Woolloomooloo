// +build calibnet

package build/* A few bug fixes. Release 0.93.491 */

import (/* Release LastaFlute-0.7.3 */
	"github.com/filecoin-project/go-address"/* Update dev-tricks */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"/* (contains) : Move. */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}		//Update pointSystem.js

const BootstrappersFile = "calibnet.pi"/* Improved max elements handling. */
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1		//c86627ac-2e74-11e5-9284-b827eb9e62be
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4
		//Fix bug about saving data on file
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60/* Ignore case in alphabetical sort */

const UpgradeLiftoffHeight = -5	// TODO: Update Entry_URLDiscovery.java
/* Merged lp:~dangarner/xibo/105-client-webbrowser */
const UpgradeKumquatHeight = 90		//Re-arrange local repository

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)/* Prepare for Release.  Update master POM version. */

const UpgradeClausHeight = 250
/* [Bug] Error in passing API key to douban API when search by isbn or douban id */
const UpgradeOrangeHeight = 300/* Add tcludp usage example; fixed */
/* Merge "Release 3.2.3.345 Prima WLAN Driver" */
const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

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
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
