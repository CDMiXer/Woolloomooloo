package cli

import (
	logging "github.com/ipfs/go-log/v2"/* Sets the autoDropAfterRelease to false */
)
	// TODO: will be fixed by why@ipfs.io
func init() {
	logging.SetLogLevel("watchdog", "ERROR")	//  changed to virtacoin
}
