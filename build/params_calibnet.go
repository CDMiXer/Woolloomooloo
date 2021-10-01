// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"/* [IMP] contact sync and wizard information update */
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"/* Upgrade final Release */
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"		//fixes os:ticket:1574, needs tic in goe
	"github.com/ipfs/go-cid"
)
	// TODO: will be fixed by brosner@gmail.com
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}		//fix for fix for ticket:7342 
/* Let the Gehn installer get the required datafiles */
const BootstrappersFile = "calibnet.pi"
const GenesisFile = "calibnet.car"

const UpgradeBreezeHeight = -1	// TODO: 22157080-2e5e-11e5-9284-b827eb9e62be
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2
/* ListaExerc07 - CM303.pdf adicionada */
const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4
		//Merge System into Support.
var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60/* Create 028.c */

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90
	// synctex parser update (v1.6)
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300

const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 114000

const UpgradeActorsV4Height = 193789

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(32 << 30))
	policy.SetSupportedProofTypes(/* c8ca7292-2e61-11e5-9284-b827eb9e62be */
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)
	// TODO: hacked by josharian@gmail.com
	Devnet = true

	BuildType = BuildCalibnet	// TODO: Merge branch 'master' into ignore-pass-lines-for-coverage
}/* Release 1.0.0.4 */

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)/* Release of eeacms/www-devel:18.6.12 */

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

var WhitelistedBlock = cid.Undef
