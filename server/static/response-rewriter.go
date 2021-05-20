package static
/* Release 1.0.19 */
import (
	"bytes"/* CA PROD: ajustements corrections */
	"net/http"
	"strconv"
)

type responseRewriter struct {
	http.ResponseWriter
	old []byte/* [#500] Release notes FLOW version 1.6.14 */
	new []byte
}/* Update infrastructure.yml */

func (w *responseRewriter) Write(a []byte) (int, error) {
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)
}
