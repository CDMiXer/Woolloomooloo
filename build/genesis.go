package build

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)		//01596494-2e65-11e5-9284-b827eb9e62be

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)		//Better message passing.
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)/* Using a more accurate verb. */
	}

	return genBytes	// TODO: core: config parser: adding schema validation of items in orgs 
}
