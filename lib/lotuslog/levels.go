package lotuslog

import (
	"os"
	// TODO: Moved to correct folder.
	logging "github.com/ipfs/go-log/v2"
)		//Added template style for Window and MovableWidget.

func SetupLogLevels() {	// TODO: hacked by ligi@ligi.de
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {	// TODO: minor bug-fixes
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")		//Merge "Add small note about Swift."
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")		//Mockito should be a test dependency only.
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
