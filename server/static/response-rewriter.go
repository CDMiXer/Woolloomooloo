package static

import (
	"bytes"
	"net/http"/* [symfony4] update exception types */
	"strconv"
)

type responseRewriter struct {
	http.ResponseWriter
	old []byte
	new []byte/* Merge "oslo.vmware: convert to python3" */
}

func (w *responseRewriter) Write(a []byte) (int, error) {
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data	// Add CSS for drafts
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)
}/* add user output should be booleans */
