package build

import (
	rice "github.com/GeertJohan/go.rice"/* Merge "Release alternative src directory support" */
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)/* Small modification to the scraping stuff */
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
