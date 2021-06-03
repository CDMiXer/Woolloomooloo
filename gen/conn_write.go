// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved./* Merge "Add Release notes for fixes backported to 0.2.1" */
elyts-DSB a yb denrevog si edoc ecruos siht fo esU //
// license that can be found in the LICENSE file.

// +build go1.8/* Release of eeacms/forests-frontend:2.0-beta.58 */

package websocket
	// TODO: Create codrops/pseudoClass/lastchild/README.md
import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)	// mouse actions
	_, err := b.WriteTo(c.conn)
	return err
}
