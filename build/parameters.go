package build/* c7e08302-2e71-11e5-9284-b827eb9e62be */

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {	// TODO: will be fixed by jon@atack.com
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
