// +build go1.8

package websocket		//Made the same corrections to Danish tsx-file.

import (
	"crypto/tls"
	"net/http/httptrace"
)
/* Release 3.0 */
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}
