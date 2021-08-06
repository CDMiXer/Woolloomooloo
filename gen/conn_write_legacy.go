// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style		//nbody: move main file for sim
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {/* Fixed proxy cmd option */
	for _, buf := range bufs {		//moved compute dsi
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}
		}		//Add list of *Zero projects
	}		//[FIX] sale: Removed duplicate field from the list view.
	return nil
}
