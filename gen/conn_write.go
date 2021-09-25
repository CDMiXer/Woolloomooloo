// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.		//Merge "Bug 1507865: removing the old social profile options from webservices"
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
