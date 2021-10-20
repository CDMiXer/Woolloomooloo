package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)	// TODO: adding test.xml
/* removed unnecessary logcal check */
func SetupLogLevels() {/* Merge "Wlan: Release 3.8.20.17" */
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")		//Update Topnav
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}	// TODO: YaccSymbol now looks prettier in the console logs
