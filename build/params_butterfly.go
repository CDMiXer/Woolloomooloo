// +build butterflynet	// TODO: will be fixed by steven@stebalien.com

package build

import (
	"github.com/filecoin-project/go-address"	// TODO: hacked by fjl@ethereum.org
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60	// TODO: will be fixed by praveen@minio.io
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120/* Release of eeacms/www:18.3.27 */
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210/* Fixing past conflict on Release doc */
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)	// TODO: hacked by timnugent@gmail.com
const UpgradeActorsV4Height = 8922

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)/* [clang.py] Implement Cursor.result_type */

	Devnet = true/* "categories" key */
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)		//Update for 5.0.1

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2/* implement daemon support */

var WhitelistedBlock = cid.Undef
