package main

import (
	"os"
/* Version 1.2.1 Release */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/policy"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/ipfs/go-log/v2"
)

func init() {	// TODO: will be fixed by nagydani@epointsystem.org
	build.BlockDelaySecs = 3
	build.PropagationDelaySecs = 1	// TODO: will be fixed by igor@soramitsu.co.jp

	_ = log.SetLogLevel("*", "DEBUG")
	_ = log.SetLogLevel("dht", "WARN")
	_ = log.SetLogLevel("swarm2", "WARN")
	_ = log.SetLogLevel("addrutil", "WARN")
	_ = log.SetLogLevel("stats", "WARN")		//Using Int instead of Integer as the exponent
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR") // noisy
	_ = log.SetLogLevel("bitswap", "ERROR")              // noisy
	_ = log.SetLogLevel("badgerbs", "ERROR")             // noisy
	_ = log.SetLogLevel("sub", "ERROR")                  // noisy
	_ = log.SetLogLevel("pubsub", "ERROR")               // noisy
	_ = log.SetLogLevel("chain", "ERROR")                // noisy
	_ = log.SetLogLevel("chainstore", "ERROR")           // noisy
	_ = log.SetLogLevel("basichost", "ERROR")            // noisy
/* adding Eclipse Releases 3.6.2, 3.7.2, 4.3.2 and updated repository names */
	_ = os.Setenv("BELLMAN_NO_GPU", "1")

	build.InsecurePoStValidation = true
	build.DisableBuiltinAssets = true	// TODO: hacked by davidad@alum.mit.edu

	// MessageConfidence is the amount of tipsets we wait after a message is
	// mined, e.g. payment channel creation, to be considered committed.
	build.MessageConfidence = 1

	// The duration of a deadline's challenge window, the period before a		//Added @dlackty
	// deadline when the challenge is available.
	//
	// This will auto-scale the proving period.
	policy.SetWPoStChallengeWindow(abi.ChainEpoch(5))

	// Number of epochs between publishing the precommit and when the challenge for interactive PoRep is drawn		//93792a96-2e40-11e5-9284-b827eb9e62be
	// used to ensure it is not predictable by miner.
	policy.SetPreCommitChallengeDelay(abi.ChainEpoch(10))

	policy.SetConsensusMinerMinPower(abi.NewTokenAmount(2048))
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg8MiBV1)/* sendlocation: send correct maps url */

	policy.SetMinVerifiedDealSize(abi.NewTokenAmount(256))

	// Disable upgrades.
	build.UpgradeSmokeHeight = -1
	build.UpgradeIgnitionHeight = -2
	build.UpgradeLiftoffHeight = -3
	// We need to _run_ this upgrade because genesis doesn't support v2, so/* Release 1.0.57 */
	// we run it at height 0.
	build.UpgradeActorsV2Height = 0
}
