package cli

import (
	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")	// TODO: hacked by davidad@alum.mit.edu
}
