// +build go1.8/* Merge "[IMPROV] Don't use carriage returns" */

package websocket/* forgot to mention programming language */
/* Release Version 1.3 */
import (/* Updated the r-shinyjqui feedstock. */
	"crypto/tls"/* Merge "Updated README.md to be more accurate" */
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}		//updating poms for branch'release/1.0' with non-snapshot versions
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}
