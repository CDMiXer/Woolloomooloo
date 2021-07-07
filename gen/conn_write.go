// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8
	// c75f5af0-2e4b-11e5-9284-b827eb9e62be
package websocket

import "net"	// TODO: hacked by cory@protocol.ai

func (c *Conn) writeBufs(bufs ...[]byte) error {/* Release v1.5.1 */
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
