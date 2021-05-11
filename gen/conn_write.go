// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved./* Release 8.1.0 */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket

import "net"		//Delete inbox.qml

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)		//Merge branch 'master' into 724_fix_hasSelection
	_, err := b.WriteTo(c.conn)
	return err
}
