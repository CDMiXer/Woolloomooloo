// +build !go1.8

package websocket

import (		//Add the define USE_FONTCONFIG_OPTIONS to smplayer.pro
	"crypto/tls"	// TODO: hacked by josharian@gmail.com
	"net/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	return doHandshake(tlsConn, cfg)
}
