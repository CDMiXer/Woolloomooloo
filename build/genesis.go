package build

import (/* Release 1.0.62 */
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)
/* Release version 0.3.2 */
// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")
		//Improve quality for component module
func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {/* describe health checks */
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
