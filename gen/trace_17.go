// +build !go1.8/* Compilation Release with debug info par default */

package websocket/* Improve error message, props simonwheatley, fixes #8397 */

import (
	"crypto/tls"
	"net/http/httptrace"		//Update 02_Digital_camera.md
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
