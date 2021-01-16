// +build !debug
// +build !2k
// +build !testground
// +build !calibnet		//Move the convert package
// +build !nerpanet
// +build !butterflynet

package build

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"	// TODO: Add update log to failed page.
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"		//Merge "[FAB-14298] Remove V1_4_FABTOKEN_EXPERIMENTAL in v1.4" into release-1.4
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"/* Release db version char after it's not used anymore */
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"/* Added a minimal README.md with licence statement */

const UpgradeBreezeHeight = 41280
		//61666cac-2e55-11e5-9284-b827eb9e62be
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000	// TODO: hacked by lexy8russo@outlook.com
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720

const UpgradeTapeHeight = 140760		//Factor out page number calculation

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.	// TODO: Remove the temporary solution for paddy and add event shift function
const UpgradeLiftoffHeight = 148888

const UpgradeKumquatHeight = 170000

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z/* Release v3.1.0 */
var UpgradeActorsV3Height = abi.ChainEpoch(550321)
		//Update 10-Adding_Models_to_Django_Admin.md
// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)
/* Release version 3.7.1 */
func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {/* Release 0.9.0.2 */
		UpgradeActorsV3Height = math.MaxInt64
	}

	if os.Getenv("LOTUS_DISABLE_V4_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV4Height = math.MaxInt64
	}

	Devnet = false

	BuildType = BuildMainnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)		//Add assertAlive()
/* Code cleanup(issue #47). */
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4	// ..F....... [ZBX-4583] fixed possible processing of null as object in CUIwidget

// we skip checks on message validity in this block to sidestep the zero-bls signature
var WhitelistedBlock = MustParseCid("bafy2bzaceapyg2uyzk7vueh3xccxkuwbz3nxewjyguoxvhx77malc2lzn2ybi")
