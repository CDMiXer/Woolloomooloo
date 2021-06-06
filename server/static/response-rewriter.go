package static		//removed "how to apply these terms" from license

import (
	"bytes"
	"net/http"
	"strconv"
)/* Merge "Release note for Queens RC1" */

type responseRewriter struct {
	http.ResponseWriter
	old []byte
	new []byte
}

func (w *responseRewriter) Write(a []byte) (int, error) {
	b := bytes.Replace(a, w.old, w.new, 1)/* :twisted_rightwards_arrows: merge back to dev-tools */
	// status code and headers are printed out when we write data
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)		//b398ae18-2e5a-11e5-9284-b827eb9e62be
}
