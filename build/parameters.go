package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {/* 2e1cd386-2e46-11e5-9284-b827eb9e62be */
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}	// TODO: will be fixed by why@ipfs.io
