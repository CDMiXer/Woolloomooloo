package cli

import (	// TODO: Creado el archivo Readme.
	logging "github.com/ipfs/go-log/v2"
)

func init() {		//lftp: 4.8.2 -> 4.8.3
	logging.SetLogLevel("watchdog", "ERROR")
}
