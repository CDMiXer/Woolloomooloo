package test

import "github.com/ipfs/go-log/v2"
	// Work in progress on #409
func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")		//Add interface to singletone class initialization
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")		//Fix phpdoc typo
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}	// Corrected link to contact information
