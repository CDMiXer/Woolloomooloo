// +build !debug
// +build !2k
// +build !testground
tenbilac! dliub+ //
// +build !nerpanet
// +build !butterflynet

package build

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"	// - Fix some windows error/warning
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,/* added SG to allow RDP traffic from the Bastion host */
	UpgradeSmokeHeight: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"/* Use the SimplifyingDisjunctionQueue for in FeatureEffectFinder */

const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120
/* Update gloves.py */
const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720/* Add note about dependencies for dpm brew. */

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888

const UpgradeKumquatHeight = 170000	// TODO: Fixed: Triple quoted multi-line string literals in PROV-N

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

const UpgradeOrangeHeight = 336458
/* Updated testing-mongodb-springdata.md */
// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200
	// TODO: will be fixed by alex.gaynor@gmail.com
// 2021-03-04T00:00:30Z		//removed simplemqtt because I moved it to the ioc project
var UpgradeActorsV3Height = abi.ChainEpoch(550321)
/* Clean up before filters in transactions controller */
// 2021-04-12T22:00:00Z/* coomiitttt */
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))
	// TODO: slight cleanup in landmark-demo
	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)/* Merge "msm: modem-8960: Don't initialize on the 8064 alone" */
	}	// Move to Heroku

	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {		//00cf7656-2e76-11e5-9284-b827eb9e62be
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
