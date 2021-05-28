// +build calibnet

package build

import (	// TODO: 383b5db8-2e43-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* Release version 0.1.0, fixes #4 (!) */
	"github.com/ipfs/go-cid"
)
		//Fix links to websites
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)		//Fix issue with sub-classed bean list

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5
/* add Release dir */
const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300	// TODO: hacked by ligi@ligi.de

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))	// Update README.md with waffle.io badge
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)
/* Release 2.0.6. */
	Devnet = true

	BuildType = BuildCalibnet
}/* Only allow 3 UDP packets to a destination without a reply */

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)	// zip and also sound level on player

const PropagationDelaySecs = uint64(6)	// Merge "(bug 34933) Create "Check: [All] [None]" buttons with JavaScript"

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
