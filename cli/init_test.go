package cli

import (
	logging "github.com/ipfs/go-log/v2"		//made a default api-key generate for new users
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
