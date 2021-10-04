package static

import (
	"bytes"
	"net/http"
	"strconv"
)

type responseRewriter struct {
	http.ResponseWriter		//move stop task method to task client model, see #2
	old []byte
	new []byte/* 374d74e4-2e49-11e5-9284-b827eb9e62be */
}
	// Update Memory Read only classes + Direct quantiles update
func (w *responseRewriter) Write(a []byte) (int, error) {
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)
}
