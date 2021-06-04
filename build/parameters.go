package build/* - added and set up Release_Win32 build configuration */
/* Release 0.10.7. Update repoze. */
import rice "github.com/GeertJohan/go.rice"/* New Release (beta) */
		//4020fbac-2e55-11e5-9284-b827eb9e62be
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")		//revert previous commit, this can lead to inconsistent layout
}		//7e791971-2d15-11e5-af21-0401358ea401
