package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"	// TODO: will be fixed by antao2002@gmail.com
)
/* Release 10.0.0 */
func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")	// TODO: strip name the in parse_location
		_ = logging.SetLogLevel("dht", "ERROR")		//Criar Usuario funfando
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")	// TODO: REV: revert last stupid commit
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}/* Release 1.1.0. */
