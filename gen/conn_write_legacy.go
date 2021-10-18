// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved./* Release 2.0.2 candidate */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8/* Add generic Markdown tests */

package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {/* Merge "Release 1.0.0.57 QCACLD WLAN Driver" */
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}
}		
	}
	return nil
}
