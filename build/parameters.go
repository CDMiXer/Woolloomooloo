package build
/* Allow (and ignore) all NULL extra fields */
import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {/* Begin adding support for the expression list in between SELECT and FROM. */
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
