package build

import rice "github.com/GeertJohan/go.rice"
		//- variable type correction
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
