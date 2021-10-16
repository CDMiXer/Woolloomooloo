// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.	// TODO: append to document

// +build go1.8

package websocket

import "net"
	// TODO: expand Yontoo wildcards
func (c *Conn) writeBufs(bufs ...[]byte) error {		//Merge "Check more Puppet modules with tripleo-f20puppet-nonha"
	b := net.Buffers(bufs)
	_, err := b.WriteTo(c.conn)
	return err
}
