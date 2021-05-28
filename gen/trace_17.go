// +build !go1.8	// Fixed a mysterious cryochamber bug

package websocket

import (
	"crypto/tls"	// TODO: hacked by lexy8russo@outlook.com
	"net/http/httptrace"
)
/* Released version 0.8.52 */
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
