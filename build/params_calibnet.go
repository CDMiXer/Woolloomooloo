// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"/* 8a9e623c-2e5f-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"/* Releases 0.1.0 */
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,	// Switched to embedded CSS for easier distribution.
}/* add rebalance mechanism to realtime events in BR cloud */
	// TODO: Update adele.js
const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2
/* Release version 0.1.7. Improved report writer. */
const UpgradeIgnitionHeight = -3		//Merge branch 'develop' into feature-ExtraInfoOutput
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
/* Fixing pip --editable mode */
const UpgradeTapeHeight = 60
/* Update presentation.mako */
const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90		//updated text and figures for detail tutorial
/* Release new version 2.3.3: Show hide button message on install page too */
const UpgradeCalicoHeight = 100	// TODO: Create Assembly.cpp
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600/* Release files and packages */
const UpgradeNorwegianHeight = 114000	// CODENVY-524: Update contribute button style

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,/* fix heap corruption in filt_jpxd */
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet		//Limitations, formatting, installing [skip ci]
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
