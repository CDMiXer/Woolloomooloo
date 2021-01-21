// +build !go1.8
		//Merge "Hygiene: Eliminate api fixmes from PageApi"
package websocket

import (
	"crypto/tls"
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
