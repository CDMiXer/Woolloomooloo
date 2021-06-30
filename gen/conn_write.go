// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.	// tried to add substitute

// +build go1.8
	// Updated: advancedrestclient 13.0.4
package websocket

import "net"/* RulesBase ajoutée */

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
