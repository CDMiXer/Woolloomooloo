// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.		//Automatic changelog generation for PR #38871 [ci skip]
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "net"	// TODO: add logging for LayoutMenu DefaultMenuAccessProvider

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
