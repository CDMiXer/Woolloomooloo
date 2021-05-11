// +build !go1.8

package websocket
/* fix packaging tag */
import (
	"crypto/tls"
	"net/http/httptrace"
)	// TODO: hacked by magik6k@gmail.com

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {	// TODO: Fixed the AMI id in the configuration file.
	return doHandshake(tlsConn, cfg)
}
