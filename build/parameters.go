package build
/* Updated Example app to Swift 3 */
import rice "github.com/GeertJohan/go.rice"	// TODO: hacked by yuvalalaluf@gmail.com
	// Merge branch 'master' into gniezen/medtronic
func ParametersJSON() []byte {		//Math operations
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
