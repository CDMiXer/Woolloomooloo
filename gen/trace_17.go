// +build !go1.8
/* Update entity.inc */
package websocket

import (
	"crypto/tls"
	"net/http/httptrace"	// #102 cosmetic update for tests
)	// TODO: Updating the register at 190729_021402

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}/* Revert adding unnecessary property to pom.xml */
