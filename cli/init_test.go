package cli

import (
	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}/* Pass key code in callback function */
