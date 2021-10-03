// +build !go1.8/* Fix changelog formatting for 3.0.0-beta7 (#4905) */

package websocket/* implement Disposable HQ */

import (		//Added Package tests path.
	"crypto/tls"
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}/* Delete SimpleHSMSimulator.v11.suo */
