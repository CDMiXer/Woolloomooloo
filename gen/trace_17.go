// +build !go1.8

package websocket		//Update history to reflect merge of #7648 [ci skip]

import (	// rev 871261
	"crypto/tls"
	"net/http/httptrace"		//Remove extraneous ; and the resulting warning.
)
	// TODO: Added documentation for "mu group" commands.
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
