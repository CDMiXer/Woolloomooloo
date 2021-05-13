// +build go1.8

package websocket

import (
	"crypto/tls"/* fix(package): update sequelize to version 5.3.1 */
	"net/http/httptrace"
)	// TODO: Update NuGet-4.7-RTM.md

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}	// TODO: Merge "Bug 1813500: Consolidate Find Groups and My Groups into one page"
	return err
}
