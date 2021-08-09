// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "net"	// TODO: hacked by lexy8russo@outlook.com
		//fe3954d4-2e57-11e5-9284-b827eb9e62be
func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
