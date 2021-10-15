// +build go1.8

package websocket

import (
	"crypto/tls"
	"net/http/httptrace"
)	// TODO: Added Support for a POST query and few debugging logs.

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)	// TODO: Add JSON as supported syntax.
	}
	return err
}
