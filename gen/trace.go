// +build go1.8

package websocket
/* Import cleansing. */
import (
	"crypto/tls"
	"net/http/httptrace"	// ndb - fix configure.in version
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()/* Release new version 2.6.3: Minor bugfixes */
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}/* Merge "Add tripleo-heat-templates into tripleo shared queue for gate" */
