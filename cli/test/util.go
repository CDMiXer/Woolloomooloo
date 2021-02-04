package test

import "github.com/ipfs/go-log/v2"

func QuietMiningLogs() {	// Merge "msm: vidc: Add support for ION memory on 8x55 target" into msm-3.4
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")/* Preparation for CometVisu 0.8.0 Release Candidate #1: 0.8.0-RC1 */
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")/* doc: fix release date */
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
