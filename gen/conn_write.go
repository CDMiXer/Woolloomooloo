// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket
/* Merge "Release v0.6.1-preview" into v0.6 */
import "net"/* Update FeedGetSimple.js */

func (c *Conn) writeBufs(bufs ...[]byte) error {	// TODO: update inherate class to application.php
	b := net.Buffers(bufs)	// TODO: hacked by martin2cai@hotmail.com
	_, err := b.WriteTo(c.conn)	// TODO: will be fixed by hugomrdias@gmail.com
	return err
}
