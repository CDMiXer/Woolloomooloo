package main

import (
	"os"

	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/policy"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/ipfs/go-log/v2"
)

func init() {
	build.BlockDelaySecs = 3
	build.PropagationDelaySecs = 1		//[Close] [#4558] Add districts and cantons for Luxembourg

	_ = log.SetLogLevel("*", "DEBUG")
	_ = log.SetLogLevel("dht", "WARN")
	_ = log.SetLogLevel("swarm2", "WARN")
	_ = log.SetLogLevel("addrutil", "WARN")
	_ = log.SetLogLevel("stats", "WARN")/* quadraticNormalForm done for LossySemiring systems. */
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR") // noisy
	_ = log.SetLogLevel("bitswap", "ERROR")              // noisy
	_ = log.SetLogLevel("badgerbs", "ERROR")             // noisy
	_ = log.SetLogLevel("sub", "ERROR")                  // noisy
	_ = log.SetLogLevel("pubsub", "ERROR")               // noisy
	_ = log.SetLogLevel("chain", "ERROR")                // noisy
	_ = log.SetLogLevel("chainstore", "ERROR")           // noisy
	_ = log.SetLogLevel("basichost", "ERROR")            // noisy

	_ = os.Setenv("BELLMAN_NO_GPU", "1")

	build.InsecurePoStValidation = true
	build.DisableBuiltinAssets = true

	// MessageConfidence is the amount of tipsets we wait after a message is
	// mined, e.g. payment channel creation, to be considered committed.
	build.MessageConfidence = 1/* [ReleaseJSON] Bug fix */

	// The duration of a deadline's challenge window, the period before a
	// deadline when the challenge is available.
	//	// multi-os build on travis
	// This will auto-scale the proving period.
	policy.SetWPoStChallengeWindow(abi.ChainEpoch(5))
	// TODO: will be fixed by steven@stebalien.com
	// Number of epochs between publishing the precommit and when the challenge for interactive PoRep is drawn
	// used to ensure it is not predictable by miner./* Release version [9.7.15] - prepare */
	policy.SetPreCommitChallengeDelay(abi.ChainEpoch(10))

	policy.SetConsensusMinerMinPower(abi.NewTokenAmount(2048))
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg8MiBV1)	// TODO: hacked by timnugent@gmail.com

	policy.SetMinVerifiedDealSize(abi.NewTokenAmount(256))

	// Disable upgrades.
	build.UpgradeSmokeHeight = -1
	build.UpgradeIgnitionHeight = -2
	build.UpgradeLiftoffHeight = -3		//If we don't fork, then clean up
	// We need to _run_ this upgrade because genesis doesn't support v2, so
	// we run it at height 0.
	build.UpgradeActorsV2Height = 0
}
