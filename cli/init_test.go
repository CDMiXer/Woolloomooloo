package cli

import (
	logging "github.com/ipfs/go-log/v2"
)
/* Fix share feed between users: test on feed URL, not page URL */
func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
