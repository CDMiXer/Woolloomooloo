// +build !debug
// +build !2k
// +build !testground
// +build !calibnet		//e45caab0-2ead-11e5-83c3-7831c1d44c14
// +build !nerpanet
// +build !butterflynet		//The same code works in Linux - so ifdefs removed

package build

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"	// Removed version from request packages

const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000/* 1.0.0 Release (!) */
const UpgradeRefuelHeight = 130800
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
const UpgradeActorsV2Height = 138720/* Release notes and JMA User Guide */
	// Fix cfy3 blueprint generation tests
const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier./* Исправлены имена переменных */
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888		//computing content-type for response not required at this time

const UpgradeKumquatHeight = 170000

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)		//correct clean targets

const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z		//color.py - clear _terminfo_params in win32 mode
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)
/* Updating docker file */
func init() {	// Don’t start disseminating when you receive a join request.
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
/* Update django-formtools from 1.0 to 2.1 */
	Devnet = false

	BuildType = BuildMainnet/* Released version 0.8.14 */
}/* allow special keys */

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

// we skip checks on message validity in this block to sidestep the zero-bls signature
var WhitelistedBlock = MustParseCid("bafy2bzaceapyg2uyzk7vueh3xccxkuwbz3nxewjyguoxvhx77malc2lzn2ybi")
