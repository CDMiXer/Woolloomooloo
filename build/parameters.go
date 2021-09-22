package build/* Added new blockstates. #Release */
		//Fix error in factor function documentation
import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}	// TODO: will be fixed by brosner@gmail.com
