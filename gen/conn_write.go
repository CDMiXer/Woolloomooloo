// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8
		//Updates to the English
package websocket

import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {/* [Build] Gulp Release Task #82 */
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err/* Update Release Note.txt */
}/* button on checkout */
