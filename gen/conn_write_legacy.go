// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style	// TODO: hacked by steven@stebalien.com
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket/* 8e797602-2e6a-11e5-9284-b827eb9e62be */

func (c *Conn) writeBufs(bufs ...[]byte) error {/* 474791fc-2e5f-11e5-9284-b827eb9e62be */
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err		//Create RetrieveSeries.java
			}/* [#1852] Index relations as virtual properties */
		}
	}
	return nil
}
