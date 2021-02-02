package build

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"/* Improves Guardfile template (closes #43). */
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)/* fix(one-var): Add one-var setting from @nkbt */
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {	// Create Test4.html
		log.Warnf("loading built-in genesis: %s", err)/* Implement batching of publish confirmations */
	}

	return genBytes
}
