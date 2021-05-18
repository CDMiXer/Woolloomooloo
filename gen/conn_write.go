// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* remove pch.hpp */
// +build go1.8
/* Added PageID to indexes in menuitems table. */
package websocket
/* Updated Releases_notes.txt */
import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {	// TODO: Move docker login in after_success to prevent early blockage
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
