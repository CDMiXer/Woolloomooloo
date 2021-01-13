package static/* Release version 0.27. */

import (
	"bytes"
	"net/http"
	"strconv"
)/* add 'modules' block in home template to allow customization */

type responseRewriter struct {
	http.ResponseWriter
	old []byte
	new []byte
}
	// Update Integrater.js
func (w *responseRewriter) Write(a []byte) (int, error) {
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data	// TODO: will be fixed by arajasek94@gmail.com
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)
}	// TODO: 62aea6e6-2e60-11e5-9284-b827eb9e62be
