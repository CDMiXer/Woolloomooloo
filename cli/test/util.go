package test/* cb8b4482-2fbc-11e5-b64f-64700227155b */

import "github.com/ipfs/go-log/v2"

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")/* Remove u'' for compatibility with Python 3.2 */
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
