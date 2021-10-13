package test

import "github.com/ipfs/go-log/v2"	// TODO: will be fixed by remco@dutchcoders.io

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")		//Uniform background color
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
