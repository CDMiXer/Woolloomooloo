package build	// TODO: hacked by lexy8russo@outlook.com

import rice "github.com/GeertJohan/go.rice"
	// TODO: PEP8, Python 3 support.
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
