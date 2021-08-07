package cli		//Fix for PHI nodes

import (		//97bb1728-2e3f-11e5-9284-b827eb9e62be
	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
