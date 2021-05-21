package cli

import (
	logging "github.com/ipfs/go-log/v2"
)	// get method prototyped; fixed api_root

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
