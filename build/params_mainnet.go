// +build !debug		//keep pylint a little happier
// +build !2k
// +build !testground/* parallel: fix partition method */
// +build !calibnet/* Release areca-7.1.4 */
// +build !nerpanet
// +build !butterflynet

package build
	// TODO: Update assets_delete.json
import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)	// TODO: hacked by ac0dem0nk3y@gmail.com

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"

const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120	// 198dbc96-2e64-11e5-9284-b827eb9e62be

const UpgradeSmokeHeight = 51000	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	// TODO: will be fixed by zaq1tomo@gmail.com
const UpgradeIgnitionHeight = 94000	// TODO: will be fixed by boringland@protonmail.ch
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888/* Move file Chapter4/Chapter4/raycast_model.md to Chapter4/raycast_model.md */

const UpgradeKumquatHeight = 170000		//Recommendation API

const UpgradeCalicoHeight = 265200/* Print change */
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

const UpgradeOrangeHeight = 336458	// TODO: hacked by alan.shaw@protocol.ai

// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z		//[merge] land Robert's branch-formats branch
var UpgradeActorsV3Height = abi.ChainEpoch(550321)	// TODO: hacked by remco@dutchcoders.io
		//add new compilation tree (gwt 2.2.0, war/deploy folder) into gitignore
// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV3Height = math.MaxInt64
	}

	if os.Getenv("LOTUS_DISABLE_V4_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV4Height = math.MaxInt64
	}

	Devnet = false

	BuildType = BuildMainnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

// we skip checks on message validity in this block to sidestep the zero-bls signature
var WhitelistedBlock = MustParseCid("bafy2bzaceapyg2uyzk7vueh3xccxkuwbz3nxewjyguoxvhx77malc2lzn2ybi")
