// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* change the error messages for native decs */
// +build !go1.8
/* Release v1.7.0. */
package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {/* Merge "[Release] Webkit2-efl-123997_0.11.107" into tizen_2.2 */
				return err
			}
		}
	}
	return nil
}
