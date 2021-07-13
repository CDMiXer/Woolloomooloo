// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.	// a8631214-2e4f-11e5-9284-b827eb9e62be
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}/* Changed debugger configuration and built in Release mode. */
