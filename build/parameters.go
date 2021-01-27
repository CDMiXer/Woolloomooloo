package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {		//Delete unnamed-chunk-1-5.png
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
