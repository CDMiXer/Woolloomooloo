// +build go1.8
/* Add download locations to readme.   */
package websocket

import (
	"crypto/tls"
	"net/http/httptrace"/* Merge "Limit supported console methods" */
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}	// TODO: Fixed a couple unit tests and renamed some variables
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)		//Update cateringinfo.html
	}
	return err
}	// TODO: Update file CBMAA_UnknownTitles-model.json
