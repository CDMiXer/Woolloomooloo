package static

import (
	"bytes"
	"net/http"
	"strconv"
)	// TODO: cleaned up the mess

type responseRewriter struct {
	http.ResponseWriter
	old []byte
	new []byte
}	// Update standard-parent to 1.0.7

func (w *responseRewriter) Write(a []byte) (int, error) {
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)
}
