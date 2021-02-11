// +build go1.8

package websocket		//Merge "msm:pm-8x60: Do not generate access flag faults for the kernel mappings"
/* Rename ReleaseNote.txt to doc/ReleaseNote.txt */
import (		//prime checking using inbuilt functions
	"crypto/tls"
	"net/http/httptrace"
)
	// Delete Administrator.xml
func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}
