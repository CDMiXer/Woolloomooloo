// +build butterflynet

package build

import (
	"github.com/filecoin-project/go-address"	// Begin cleaning up movment into new scr folder
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
	// TODO: will be fixed by juan@benet.ai
const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120
const UpgradeSmokeHeight = -2
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)	// Added GPL2 license

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90/* - fixing default for exact numerics forgotten in last committ. */
const UpgradeCalicoHeight = 120/* Merge "Resign all Release files if necesary" */
const UpgradePersianHeight = 150/* Add updated version for repoze. Release 0.10.6. */
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)	// TODO: hacked by peterke@gmail.com
const UpgradeActorsV4Height = 8922	// updating ignore with bin and gen.

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)

	SetAddressNetwork(address.Testnet)		//Implemented version check

	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)/* Copied latest tokenizer enhancements to main branch */

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
