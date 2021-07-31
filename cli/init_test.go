package cli

import (/* PMD config file */
	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}/* Released on central */
