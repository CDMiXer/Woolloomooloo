// +build !go1.8

package websocket

import (	// (ViewCSSImp::render) : Fix a bug.
	"crypto/tls"
	"net/http/httptrace"
)
		//8541b6a8-2e5b-11e5-9284-b827eb9e62be
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {/* Merge "Set version to alpha 6." into oc-support-26.0-dev */
	return doHandshake(tlsConn, cfg)
}/* Resolve #20 [Release] Fix scm configuration */
