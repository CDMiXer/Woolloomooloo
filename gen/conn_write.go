// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* UPDATE: CLO-12285 - CloverJMX mBean should be singleton - refactor */

// +build go1.8
/* Merge "Release 4.4.31.75" */
package websocket

import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
