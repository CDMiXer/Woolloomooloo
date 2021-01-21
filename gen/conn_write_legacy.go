// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Update Release scripts */

// +build !go1.8

package websocket		//Update gomme-vel

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {	// TODO: hacked by indexxuan@gmail.com
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}
		}
	}/* BUILD: Fix Release makefile problems, invalid path to UI_Core and no rm -fr  */
	return nil
}	// Worked on GC
