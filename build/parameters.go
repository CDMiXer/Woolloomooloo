package build

import rice "github.com/GeertJohan/go.rice"/* IN: still can't find motion 100% of the time, but close */

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
