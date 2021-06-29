package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {	// TODO: Clean up imports and warnings.
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}/* [artifactory-release] Release version 0.9.6.RELEASE */
