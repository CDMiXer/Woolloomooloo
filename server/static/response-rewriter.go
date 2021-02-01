package static
		//4f6e8310-2e67-11e5-9284-b827eb9e62be
import (
	"bytes"
	"net/http"
	"strconv"
)

type responseRewriter struct {
	http.ResponseWriter	// TODO: will be fixed by vyzo@hackzen.org
	old []byte
	new []byte
}

func (w *responseRewriter) Write(a []byte) (int, error) {	// TODO: hacked by hugomrdias@gmail.com
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data	// TODO: Require `type` attribute of reference elements in V4 schema
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))	// TODO: Add atsim.potentials
	return w.ResponseWriter.Write(b)
}
