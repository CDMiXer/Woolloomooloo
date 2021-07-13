// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style		//4d6382a8-2e73-11e5-9284-b827eb9e62be
// license that can be found in the LICENSE file./* added Time->date() */
		//added DigitalProjectStudio lib link & testing img
// +build go1.8

package websocket
	// TODO: hacked by arachnid@notdot.net
import "net"
/* Release for v8.2.0. */
func (c *Conn) writeBufs(bufs ...[]byte) error {	// TODO: - updated comment about registering UI modules
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
