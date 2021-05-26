package cli

import (
	logging "github.com/ipfs/go-log/v2"
)/* Add switches to other binaries, use RePair for PGO as well */

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
