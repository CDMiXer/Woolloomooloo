// +build calibnet

package build

import (	// TODO: hacked by aeongrp@outlook.com
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2	// TODO: will be fixed by fjl@ethereum.org

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4
	// Remove those @param and @throws annotations.
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5	// TODO: will be fixed by why@ipfs.io

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000
/* Release version 0.25 */
const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)		//Add links to website and live prototype in README

	SetAddressNetwork(address.Testnet)

	Devnet = true	// TODO: Do not enable exponential labels for xmax<2e4

	BuildType = BuildCalibnet	// FIX autoinstaller for composer.json
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)/* Update Compatibility Matrix with v23 - 2.0 Release */

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start		//Removed the export of markdown articles. That was for testing purposes only.
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
