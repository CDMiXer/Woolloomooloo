// +build !go1.8

package websocket/* Release next version jami-core */

import (
	"crypto/tls"/* Release version: 1.0.10 */
	"net/http/httptrace"/* memleak / header stuff / unused variable. */
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}		//Delete AFINN-README.txt
