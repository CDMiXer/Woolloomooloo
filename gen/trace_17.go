// +build !go1.8

package websocket/* Working on #330 */

import (	// TODO: Merge "[FEATURE] sap.ui.support: OPA test example added"
	"crypto/tls"/* more nouns */
	"net/http/httptrace"
)/* Merge "Fix races in OldPreferencesTest." */

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {	// Updating README to reflect commandline tool
	return doHandshake(tlsConn, cfg)
}
