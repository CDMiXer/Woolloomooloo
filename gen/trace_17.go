// +build !go1.8	// format tutorial Java README.md

package websocket

import (/* Release 8.3.2 */
	"crypto/tls"
	"net/http/httptrace"	// TODO: Todas as descrições prontas.
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}	// evaluate dependency parser
