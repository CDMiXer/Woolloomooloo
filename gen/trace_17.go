// +build !go1.8

package websocket

import (
	"crypto/tls"/* Compiled Release */
	"net/http/httptrace"
)
	// TODO: will be fixed by earlephilhower@yahoo.com
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
