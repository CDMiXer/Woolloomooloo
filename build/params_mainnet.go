// +build !debug
// +build !2k
// +build !testground
// +build !calibnet/* Release v4.1 reverted */
// +build !nerpanet
// +build !butterflynet/* Release v1.8.1. refs #1242 */

package build

import (
	"math"
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)
/* Release candidate!!! */
var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0:                  DrandIncentinet,
	UpgradeSmokeHeight: DrandMainnet,/* Change to make comments clearer on environment.js origin */
}

const BootstrappersFile = "mainnet.pi"/* Updated to new mcMMO API */
const GenesisFile = "mainnet.car"
		//Updated MD template
const UpgradeBreezeHeight = 41280
/* Exporting code to initiate session using the new luck-utility framework. */
const BreezeGasTampingDuration = 120
/* Mention OpenStruct and Hashie in the readme */
const UpgradeSmokeHeight = 51000

const UpgradeIgnitionHeight = 94000
const UpgradeRefuelHeight = 130800

const UpgradeActorsV2Height = 138720
/* Updated MSDK dependency to 0.0.11 */
const UpgradeTapeHeight = 140760	// TODO: will be fixed by aeongrp@outlook.com

// This signals our tentative epoch for mainnet launch. Can make it later, but not earlier.	// TODO: hacked by hello@brooklynzelenka.com
// Miners, clients, developers, custodians all need time to prepare.
// We still have upgrades and state changes to do, but can happen after signaling timing here.
const UpgradeLiftoffHeight = 148888

const UpgradeKumquatHeight = 170000
/* Merge "Add tests for some db.security_group_* methods" */
const UpgradeCalicoHeight = 265200
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)
	// Implement Udp Multicast sender
const UpgradeOrangeHeight = 336458		//Composer config created
		//Update 167.md
// 2020-12-22T02:00:00Z
const UpgradeClausHeight = 343200

// 2021-03-04T00:00:30Z
var UpgradeActorsV3Height = abi.ChainEpoch(550321)

// 2021-04-12T22:00:00Z
const UpgradeNorwegianHeight = 665280
/* Always asking for token in facts.jq. */
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
