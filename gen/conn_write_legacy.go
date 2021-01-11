// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style	// TODO: will be fixed by juan@benet.ai
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {		//add information about docs in readme
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}
		}
	}
	return nil	// TODO: disable auto-install of MarGo
}	// TODO: Create search_and_purge_app.sh
