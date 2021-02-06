package build
		//added "ASAP to SPQR"
import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {/* srcp: re-connect command port */
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")	// TODO: Mise àj our du menu
}
