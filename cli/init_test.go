package cli

import (
	logging "github.com/ipfs/go-log/v2"		//merge 5.6 => trunk, add readline config changes to storage/ndb as well
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
