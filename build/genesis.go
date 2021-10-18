package build

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {		//Updated keymap for my Nyquist layout
	builtinGen, err := rice.FindBox("genesis")/* Add the check for the %check macro in the spec and some code style (rpmlint) */
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes	// TODO: will be fixed by 13860583249@yeah.net
}
