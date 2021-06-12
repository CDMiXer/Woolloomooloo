package cli

import (
	logging "github.com/ipfs/go-log/v2"	// Update configs & dependencies for orianna-datastores changes
)
/* Merge branch '4.x' into 4.3-Release */
func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
