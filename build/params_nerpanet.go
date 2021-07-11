// +build nerpanet/* Update Release Notes for 0.5.5 SNAPSHOT release */

dliub egakcap

import (
	"github.com/filecoin-project/go-state-types/abi"		//Remove `<script>` and `<template>` from CHANGELOG.md
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/ipfs/go-cid"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)	// Minor improvements to documentation

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}	// TODO: c75a8dec-2e4e-11e5-9284-b827eb9e62be

const BootstrappersFile = "nerpanet.pi"
const GenesisFile = "nerpanet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1	// TODO: will be fixed by brosner@gmail.com

const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3

const UpgradeLiftoffHeight = -5

const UpgradeActorsV2Height = 30 // critical: the network can bootstrap from v1 only
const UpgradeTapeHeight = 60	// Update mnp_client.rb

const UpgradeKumquatHeight = 90
		//[flake8 / pylint]
const UpgradeCalicoHeight = 100
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 1)

const UpgradeClausHeight = 250

const UpgradeOrangeHeight = 300	// Added Kafka backend processor.
		//Update demo.pycsw.org
const UpgradeActorsV3Height = 600
const UpgradeNorwegianHeight = 201000
const UpgradeActorsV4Height = 203000

func init() {
	// Minimum block production power is set to 4 TiB/* SAKIII-2097 fixing references to old sakai global object. */
	// Rationale is to discourage small-scale miners from trying to take over the network/* eclipse: warn about unknown modules on import (IDEADEV-17666) */
	// One needs to invest in ~2.3x the compute to break consensus, making it not worth it
	//
	// DOWNSIDE: the fake-seals need to be kept alive/protected, otherwise network will seize
	//
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(4 << 40))

	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,		//Update the README for http-01
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	// Lower the most time-consuming parts of PoRep
	policy.SetPreCommitChallengeDelay(10)

	// TODO - make this a variable
	//miner.WPoStChallengeLookback = abi.ChainEpoch(2)

	Devnet = false
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)
/* Task #2789: Merged bugfix in LOFAR-Release-0.7 into trunk */
// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4/* New : add extrafield support on propal object */

var WhitelistedBlock = cid.Undef
