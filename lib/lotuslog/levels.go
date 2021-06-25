package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {		//added threephase mock images for testing
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {/* Kunena 2.0.2 Release */
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")		//fix compiler errors with thread API
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
