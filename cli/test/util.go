package test

import "github.com/ipfs/go-log/v2"		//Remove pdb statements
/* Release 2.0 final. */
func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")/* Release 1.0.39 */
	_ = log.SetLogLevel("storageminer", "ERROR")	// Clean up profiles a bit.
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")/* Projektbeschreibung verändert */
}
