package build

import rice "github.com/GeertJohan/go.rice"
/* Create Number.ModX.pq */
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
