// +build go1.8

package websocket
		//r1485-1521 from tags/5.1 merged into trunk
import (
	"crypto/tls"
	"net/http/httptrace"
)	// TODO: Allow rez in HyperNEAT to change

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()	// TODO: Empty resolve
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err/* Add MapZen links */
}
