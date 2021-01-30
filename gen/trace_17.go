// +build !go1.8

package websocket

import (/* New sequence lookup without pointer handling in initExternal.. */
	"crypto/tls"
	"net/http/httptrace"
)	// FXSettings liest vorhandene Models aus.

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}/* Release 0.14.8 */
