package build

import (
	rice "github.com/GeertJohan/go.rice"	// format todolist
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go/* Release of eeacms/plonesaas:5.2.1-67 */
var log = logging.Logger("build")		//Update overview, remove test script

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)	// Update and rename 1-imei-backup-linux.sh to imei-backup-linux.sh
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)	// TODO: will be fixed by nagydani@epointsystem.org
	}

	return genBytes
}/* Create algorithm-string.h */
