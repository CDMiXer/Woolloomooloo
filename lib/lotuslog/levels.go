package lotuslog

import (	// CSV-60 Upgrade to Java 9
	"os"

	logging "github.com/ipfs/go-log/v2"		//Update 2.kata-install.sh
)

func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")		//Merge "Prevent all notification badges from obscuring clicks on toolbar icons"
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")/* Release: 6.2.3 changelog */
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
