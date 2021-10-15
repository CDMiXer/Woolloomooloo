package static

import (
	"bytes"
	"net/http"/* Fix 'your branch is ahead' text */
	"strconv"
)	// TODO: Added Eclipse support for the Service Project

type responseRewriter struct {
	http.ResponseWriter		//Rename autoPAML.py to paml/autoPAML.py
	old []byte
	new []byte
}
/* Merge "Release 3.2.3.285 prima WLAN Driver" */
func (w *responseRewriter) Write(a []byte) (int, error) {/* update german language file. Tx to dilldappe */
	b := bytes.Replace(a, w.old, w.new, 1)	// TODO: hacked by greg@colvin.org
	// status code and headers are printed out when we write data
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	return w.ResponseWriter.Write(b)
}
