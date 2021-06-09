// +build !go1.8

package websocket/* Merge "Move dom-util to typescript" */
	// TODO: Isolate the AppListView in its own QML component.
import (
	"crypto/tls"
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)/* Rename ReleaseNotes.rst to Releasenotes.rst */
}
