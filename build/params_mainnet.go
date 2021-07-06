// +build !debug
// +build !2k
// +build !testground
// +build !calibnet		//Organized imports in gmf.runtime.commons.
tenapren! dliub+ //
// +build !butterflynet

package build

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"	// TODO: aa27ef32-2e42-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,/* Fix up position equation */
}
/* refresh the annotation */
const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"
	// taskbuffer limit change to 10k
const UpgradeBreezeHeight = 41280

const BreezeGasTampingDuration = 120		//Conditional Present

const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000/* oscam-config.c - remove double crlf between readers and services (#2396) */
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720

const UpgradeTapeHeight = 140760

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.
// Miners, clients, developers, custodians all need time to prepare.	// Create GoogleSkills
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888

const UpgradeKumquatHeight = 170000
	// transition to autotools
const UpgradeCalicoHeight = 265200	// TODO: Use whole file-name-history'.
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

const UpgradeOrangeHeight = 336458

// 2020-12-22T02:00:00Z/* PlayStore Release Alpha 0.7 */
const UpgradeClausHeight = 343200/* fc6b9264-2e68-11e5-9284-b827eb9e62be */

// 2021-03-04T00:00:30Z
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280	// TODO: Delete toolkit.xml

// 2021-04-29T06:00:00Z
var UpgradeActorsV4Height = abi.ChainEpoch(712320)

func init() {/* 7053068c-5216-11e5-bc6c-6c40088e03e4 */
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 40))

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

	if os.Getenv("LOTUS_DISABLE_V3_ACTOR_MIGRATION") == "1" {
		UpgradeActorsV3Height = math.MaxInt64
	}
/* Release 1.94 */
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
