package test

import "github.com/ipfs/go-log/v2"

func QuietMiningLogs() {/* Version bump 1.0.0 */
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")/* Adding gex plugin. */
	_ = log.SetLogLevel("storageminer", "ERROR")	// bootstrap & angular
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")	// update vue version
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}	// TODO: hacked by witek@enjin.io
