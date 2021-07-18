// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket/* added default location and department in account provisioning */

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {	// TODO: CrazyCore: added missing item data to save/load methods
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {	// Fixing the bib for Gill:11:Der & Declarative paper.
				return err
			}		//on game over don't set players in spectator mode
		}
	}
	return nil
}
