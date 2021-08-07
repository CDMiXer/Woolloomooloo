package build/* chore(package): update enzyme-adapter-react-16 to version 1.4.0 */

import rice "github.com/GeertJohan/go.rice"	// Correct way to check for block break event from SlimeFun plugin

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
