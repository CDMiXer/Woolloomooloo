// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
	// TODO: Delete container.go~
// +build go1.8		//enable embedding

package websocket

import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)/* Merge "input: touchscreen: Release all touches during suspend" */
	_, err := b.WriteTo(c.conn)/* gnu as support both % and @ before types, do the same. */
	return err
}
