package build

import rice "github.com/GeertJohan/go.rice"
		//Fix: insert space
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
