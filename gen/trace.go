// +build go1.8

package websocket	// TODO: will be fixed by mikeal.rogers@gmail.com
		//phpdoc fix for executable config
import (
	"crypto/tls"
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)	// TODO: hacked by 13860583249@yeah.net
	}
	return err
}
