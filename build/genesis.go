package build/* Create Release-Notes.md */

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)
		//dba34d: fix for assertion from comphelper
// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}/* Maj de l'install de Commons */

	return genBytes
}
