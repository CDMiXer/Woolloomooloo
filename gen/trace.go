// +build go1.8

package websocket

import (/* Merge "[INTERNAL] Release notes for version 1.32.0" */
	"crypto/tls"
	"net/http/httptrace"
)/* [Release 0.8.2] Update change log */
/* Release 1-136. */
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)/* Release v3.7.0 */
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err	// TODO: Update and rename Documentacion to Documentacion/Stakeholders-final
}
