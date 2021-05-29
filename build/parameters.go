package build
/* Release 1.5.0. */
import rice "github.com/GeertJohan/go.rice"/* Release version [10.6.3] - alfter build */

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")/* Release of eeacms/ims-frontend:0.6.0 */
}	// TODO: [tests/tset_si.c] Added mpfr_get_{si,ui} tests, including flags.
