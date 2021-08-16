// +build calibnet

package build		//Update CadWebsite.css

import (
	"github.com/filecoin-project/go-address"	// provide type and domainType
	"github.com/filecoin-project/go-state-types/abi"		//build: update source-map-support to version 0.5.10
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{	// TODO: Fixed bug introduced yesterday, showed "INT-Store" Project
	0: DrandMainnet,		//-Change: Renamed voxel_map.* to map.* files.
}

const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1/* Release of eeacms/ims-frontend:0.3.5 */
const BreezeGasTampingDuration = 120/* Release note additions */

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)
	// TODO: add additional remove operations to IdentitySet/IdentityMap
const UpgradeTapeHeight = 60/* Add Release Note for 1.0.5. */

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90	// TODO: will be fixed by xiemengjun@gmail.com
	// Spike: Add redirect to google from why-i-changed-careers
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250		//allowing /16
	// TODO: always include GLIB flags because they might be necessary for GST
const UpgradeOrangeHeight = 300/* Hardcode calendar manipulation */

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789	// Using if instead of while for returning single records.

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
