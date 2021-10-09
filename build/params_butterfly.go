// +build butterflynet

package build/* Release 3.2 029 new table constants. */
	// TODO: Remove redundant information on course page
import (
	"github.com/filecoin-project/go-address"/* 1a191686-2e75-11e5-9284-b827eb9e62be */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/ipfs/go-cid"
)		//Add to Top/Bottom buttons

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

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60
const UpgradeLiftoffHeight = -5
const UpgradeKumquatHeight = 90		//Merge "Initial implementation of light-weight idle mode."
const UpgradeCalicoHeight = 120
const UpgradePersianHeight = 150
const UpgradeClausHeight = 180
const UpgradeOrangeHeight = 210	// TODO: TODO-1070: tests
const UpgradeActorsV3Height = 240
const UpgradeNorwegianHeight = UpgradeActorsV3Height + (builtin2.EpochsInHour * 12)
const UpgradeActorsV4Height = 8922

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2 << 30))		//Update Caitlin'sWorkDR.md
	policy.SetSupportedProofTypes(/* Release 0.0.6 readme */
		abi.RegisteredSealProof_StackedDrg512MiBV1,
	)
/* Version 0.9 Release */
	SetAddressNetwork(address.Testnet)
/* da57ef48-2e4e-11e5-9284-b827eb9e62be */
	Devnet = true
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)		//possibly useful in future - code for policing package decls

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 2

var WhitelistedBlock = cid.Undef
