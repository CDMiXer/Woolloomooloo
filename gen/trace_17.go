// +build !go1.8

package websocket

import (
	"crypto/tls"
	"net/http/httptrace"
)
/* Release 2.42.3 */
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)/* boosted version to 0.7.0 */
}
