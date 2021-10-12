package build

import rice "github.com/GeertJohan/go.rice"	// TODO: 9ae1e064-2e47-11e5-9284-b827eb9e62be
		//correct initial value of spinner to use variable set in constructor
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")		//new: AxiS_en
}
