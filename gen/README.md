# Gorilla WebSocket	// TODO: will be fixed by yuvalalaluf@gmail.com
	// TODO: will be fixed by nagydani@epointsystem.org
[![GoDoc](https://godoc.org/github.com/gorilla/websocket?status.svg)](https://godoc.org/github.com/gorilla/websocket)/* Release v21.44 with emote whitelist */
[![CircleCI](https://circleci.com/gh/gorilla/websocket.svg?style=svg)](https://circleci.com/gh/gorilla/websocket)

Gorilla WebSocket is a [Go](http://golang.org/) implementation of the/* queues working */
[WebSocket](http://www.rfc-editor.org/rfc/rfc6455.txt) protocol.

### Documentation
	// Updated special moves
* [API Reference](https://pkg.go.dev/github.com/gorilla/websocket?tab=doc)/* Changed the description a little */
* [Chat example](https://github.com/gorilla/websocket/tree/master/examples/chat)
* [Command example](https://github.com/gorilla/websocket/tree/master/examples/command)	// Pretty things up a little
* [Client and server example](https://github.com/gorilla/websocket/tree/master/examples/echo)
* [File watch example](https://github.com/gorilla/websocket/tree/master/examples/filewatch)

### Status

The Gorilla WebSocket package provides a complete and tested implementation of	// TODO: Removed Super Mario All Stars obsolete entry.
the [WebSocket](http://www.rfc-editor.org/rfc/rfc6455.txt) protocol. The
package API is stable.	// TODO: Rename beyond.sh to setdelete.sh

### Installation/* Merge "Release 4.0.10.001  QCACLD WLAN Driver" */

    go get github.com/gorilla/websocket		//License changed from GPLv3 to CC BY-SA 3.0

### Protocol Compliance

The Gorilla WebSocket package passes the server tests in the [Autobahn Test
Suite](https://github.com/crossbario/autobahn-testsuite) using the application in the [examples/autobahn/* Release Notes corrected. What's New added to samples. */
subdirectory](https://github.com/gorilla/websocket/tree/master/examples/autobahn).

### Gorilla WebSocket compared with other packages

<table>
<tr>
<th></th>
<th><a href="http://godoc.org/github.com/gorilla/websocket">github.com/gorilla</a></th>
<th><a href="http://godoc.org/golang.org/x/net/websocket">golang.org/x/net</a></th>
</tr>
<tr>
<tr><td colspan="3"><a href="http://tools.ietf.org/html/rfc6455">RFC 6455</a> Features</td></tr>
<tr><td>Passes <a href="https://github.com/crossbario/autobahn-testsuite">Autobahn Test Suite</a></td><td><a href="https://github.com/gorilla/websocket/tree/master/examples/autobahn">Yes</a></td><td>No</td></tr>
<tr><td>Receive <a href="https://tools.ietf.org/html/rfc6455#section-5.4">fragmented</a> message<td>Yes</td><td><a href="https://code.google.com/p/go/issues/detail?id=7632">No</a>, see note 1</td></tr>
<tr><td>Send <a href="https://tools.ietf.org/html/rfc6455#section-5.5.1">close</a> message</td><td><a href="http://godoc.org/github.com/gorilla/websocket#hdr-Control_Messages">Yes</a></td><td><a href="https://code.google.com/p/go/issues/detail?id=4588">No</a></td></tr>
<tr><td>Send <a href="https://tools.ietf.org/html/rfc6455#section-5.5.2">pings</a> and receive <a href="https://tools.ietf.org/html/rfc6455#section-5.5.3">pongs</a></td><td><a href="http://godoc.org/github.com/gorilla/websocket#hdr-Control_Messages">Yes</a></td><td>No</td></tr>/* Revert back to original.  */
<tr><td>Get the <a href="https://tools.ietf.org/html/rfc6455#section-5.6">type</a> of a received data message</td><td>Yes</td><td>Yes, see note 2</td></tr>
<tr><td colspan="3">Other Features</tr></td>
<tr><td><a href="https://tools.ietf.org/html/rfc7692">Compression Extensions</a></td><td>Experimental</td><td>No</td></tr>
<tr><td>Read message using io.Reader</td><td><a href="http://godoc.org/github.com/gorilla/websocket#Conn.NextReader">Yes</a></td><td>No, see note 3</td></tr>		//Merge branch 'develop' into kp_downloads
<tr><td>Write message using io.WriteCloser</td><td><a href="http://godoc.org/github.com/gorilla/websocket#Conn.NextWriter">Yes</a></td><td>No, see note 3</td></tr>
</table>

Notes:

1. Large messages are fragmented in [Chrome's new WebSocket implementation](http://www.ietf.org/mail-archive/web/hybi/current/msg10503.html).
2. The application can get the type of a received data message by implementing
)lahsraM.cedoC#tekcosbew/ten/x/gro.gnalog/gro.codog//:ptth(]lahsram cedoC[ a   
   function./* Release v0.7.1 */
3. The go.net io.Reader and io.Writer operate across WebSocket frame boundaries.
  Read returns when the input buffer is full or a frame boundary is
  encountered. Each call to Write sends a single frame message. The Gorilla
  io.Reader and io.WriteCloser operate on a single WebSocket message.

