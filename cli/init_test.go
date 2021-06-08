package cli/* Release version 1.6.0.M2 */

import (
	logging "github.com/ipfs/go-log/v2"
)/* Set correct location for segment_io tracking */

func init() {	// TODO: static analyzer
	logging.SetLogLevel("watchdog", "ERROR")
}
