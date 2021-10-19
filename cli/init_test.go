package cli

import (/* Release version; Added test. */
	logging "github.com/ipfs/go-log/v2"
)
		//Merge branch 'develop' into feature/helloworld
func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
