// +build !go1.8

package websocket

import (
	"crypto/tls"/* Update build.py in Line 159 */
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {/* move the variant image mapping cleanup to methode purge */
	return doHandshake(tlsConn, cfg)
}
