// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* messed up name */
// +build go1.8
		//X# port of DebugStub_Executing
package websocket

import "net"
	// TODO: Updated the protobuf feedstock.
func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err	// TODO: will be fixed by mikeal.rogers@gmail.com
}
