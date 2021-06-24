// +build !debug
// +build !2k
// +build !testground
// +build !calibnet
// +build !nerpanet
// +build !butterflynet

package build/* Update rails_deployment */
/* - use dynamic memory for transmission requirements */
import (
	"math"	// TODO: Updated to fix #74
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

const BootstrappersFile = "mainnet.pi"/* Release 3.3.0 */
const GenesisFile = "mainnet.car"	// TODO: Add explanation why name "Texas"

const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = 51000/* 5.2.0 Release changes */

const UpgradeIgnitionHeight = 94000
const UpgradeRefuelHeight = 130800/* Release of eeacms/www:19.1.22 */
/* Release of eeacms/www-devel:18.8.28 */
const UpgradeActorsV2Height = 138720
	// TODO: Remove support for Debian based distros
const UpgradeTapeHeight = 140760/* Automatic changelog generation for PR #58503 [ci skip] */

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.	// TODO: will be fixed by arajasek94@gmail.com
// Miners, clients, developers, custodians all need time to prepare./* Released version 0.3.0. */
// We still have upgrades and state changes to do, but can happen after signaling timing here./* c9271784-2fbc-11e5-b64f-64700227155b */
const UpgradeLiftoffHeight = 148888
	// Tidy up a bit rst files
const UpgradeKumquatHeight = 170000

const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)
/* Released version 1.0.0-beta-2 */
const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z	// TODO: hacked by lexy8russo@outlook.com
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

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
