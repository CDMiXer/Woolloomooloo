// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.	// TODO: hacked by aeongrp@outlook.com
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
	// TODO: hacked by zaq1tomo@gmail.com
// +build go1.8	// TODO: `cyclic` rewording suggested by @kwmsmith.

package websocket

import "net"
/* Sorting ARM Sources alphabetically */
func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)/* add charset to redirect page */
	return err
}
