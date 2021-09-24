// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8
		//Semantic-UI-LESS compilation using less4j-1.13.0
package websocket/* Set text position actions */

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}/* Remove ccat tap */
		}
	}
	return nil		//Merge "Docs: Watch face performance update" into mnc-io-docs
}
