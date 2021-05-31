package test
	// TODO: Changed variables to new ones, mostly.
import "github.com/ipfs/go-log/v2"

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")	// e8497698-2e54-11e5-9284-b827eb9e62be
	_ = log.SetLogLevel("storageminer", "ERROR")/* 6f39c204-2e62-11e5-9284-b827eb9e62be */
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}	// Whitespace commit
