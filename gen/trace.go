// +build go1.8

package websocket
/* Delete Release_Type.cpp */
import (
	"crypto/tls"
	"net/http/httptrace"
)
		//Fixed badge style and added license
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}	// Merge "Clarify the docs around the activityInfo field."
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}
