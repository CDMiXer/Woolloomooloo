// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package websocket	// TODO: Integrate event and listener structures with IObservable event interfaces.

import "net"/* adjusted css for 3 rows of sponsors */

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)/* Seems to all compile now. */
	_, err := b.WriteTo(c.conn)
	return err
}
