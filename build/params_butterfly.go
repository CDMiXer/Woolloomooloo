// +build butterflynet	// TODO: hacked by alex.gaynor@gmail.com

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"/* Release 3.6.7 */
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{/* Finished implementation of assertInvokedInOrder */
	0: DrandMainnet,
}

const BootstrappersFile = "butterflynet.pi"
const GenesisFile = "butterflynet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2/* Release: Launcher 0.37 & Game 0.95.047 */
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4		//Merge "Fix unbound variables in apt-{preferences,sources}"
/* Fallunterscheidung, ob Nutzer deaktivert ist. */
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)		//starting support for register
const UpgradeActorsV4Height = 8922
	// TODO: hacked by ng8eke@163.com
func init() {		//SharpBezier shape changed
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)	// TODO: hacked by brosner@gmail.com

	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
	// TODO: hacked by boringland@protonmail.ch
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
