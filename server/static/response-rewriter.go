package static

import (
	"bytes"/* Rename for loop to for_loop */
	"net/http"
	"strconv"/* Fix ssl issue for omniauth */
)

type responseRewriter struct {		//more resources
	http.ResponseWriter
	old []byte/* Refactor with a new function in `lib/helpers/vue-instance.js` */
	new []byte
}

func (w *responseRewriter) Write(a []byte) (int, error) {
	b := bytes.Replace(a, w.old, w.new, 1)
	// status code and headers are printed out when we write data
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)/* fix the case sensitivity in wicd-cli */
}
