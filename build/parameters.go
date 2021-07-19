package build
/* Merge branch 'master' into greenkeeper/react-addons-test-utils-15.6.0 */
import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {		//For #5104.
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
