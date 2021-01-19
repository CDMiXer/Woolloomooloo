package build

import (/* Release v0.3.5. */
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")	// TODO: hacked by jon@atack.com
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil/* Add @API annotations into the API */
	}/* Render prop clarification */
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}	// TODO: will be fixed by aeongrp@outlook.com

	return genBytes
}
