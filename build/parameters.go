package build
/* do not show busy indicator in some situations */
import rice "github.com/GeertJohan/go.rice"/* Release areca-7.4.6 */

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}/* Rename index.html to Version1.0/index.html */
