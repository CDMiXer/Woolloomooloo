package test

import "github.com/ipfs/go-log/v2"
/* Release of version 2.3.1 */
func QuietMiningLogs() {		//Expand example HTML a little
	_ = log.SetLogLevel("miner", "ERROR")/* Create project.html.twig */
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
