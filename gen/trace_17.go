// +build !go1.8

package websocket/* [skip ci] Fix missing punctuation mark and formatting */

import (
	"crypto/tls"/* Renaming and deleting terminologies */
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
