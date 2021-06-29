// +build go1.8

package websocket

import (	// TODO: will be fixed by sjors@sprovoost.nl
	"crypto/tls"
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {/* Release 0.94.320 */
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()/* fix: fetch itunes EP, Single tag and remove it */
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {	// job #8321 Small addition in proxy removal section
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}		//updated the load top pos
