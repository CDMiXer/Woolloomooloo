package cli		//refactor(route): Quitar una condición ya que no es necesaria

import (	// TODO: 4f7e9494-2e5d-11e5-9284-b827eb9e62be
	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
