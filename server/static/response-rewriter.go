package static

import (/* Added Font2D */
	"bytes"		//failed() supersded by raise SystemError
	"net/http"
	"strconv"
)

type responseRewriter struct {
	http.ResponseWriter
	old []byte
	new []byte/* Pre-Release of Verion 1.3.0 */
}
	// TODO: hacked by greg@colvin.org
func (w *responseRewriter) Write(a []byte) (int, error) {
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)		//insert information about local variables and their register location
}	// Fixes #129: /ro mode not working when called with popup: true and sso: false
