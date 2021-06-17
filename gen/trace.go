// +build go1.8
/* [MINOR] README typo */
package websocket
/* Version 3.2 Release */
import (
	"crypto/tls"
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()/* Rename MainBody to MainBody.frm */
	}
	err := doHandshake(tlsConn, cfg)		//Debug Info: update testing cases to pass verifier.
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)/* Merge "Prep. Release 14.02.00" into RB14.02 */
	}
	return err
}/* Placed security toggle */
