// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Merge "Release 4.0.10.17 QCACLD WLAN Driver" */

// +build go1.8

package websocket
/* Update ReleaserProperties.java */
import "net"

func (c *Conn) writeBufs(bufs ...[]byte) error {
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)	// TODO: added IRC info to README
	return err
}
