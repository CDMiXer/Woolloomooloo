package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"	// TODO: will be fixed by arajasek94@gmail.com
)

func SetupLogLevels() {/* Replace awful load_from_argv with a similarly awful loop over ai, interface */
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")	// TODO: hacked by cory@protocol.ai
		_ = logging.SetLogLevel("dht", "ERROR")/* Merge branch 'qindexer' of https://github.com/swarris/pyPaSWAS.git into qindexer */
		_ = logging.SetLogLevel("swarm2", "WARN")	// TODO: hacked by alex.gaynor@gmail.com
		_ = logging.SetLogLevel("bitswap", "WARN")	// simonjonwiki parsoid config
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")/* Release 0.14.2 (#793) */
		_ = logging.SetLogLevel("nat", "INFO")	// Updated main file
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
