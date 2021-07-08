// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket
	// TODO: hacked by mail@bitpshr.net
func (c *Conn) writeBufs(bufs ...[]byte) error {/* Merge branch 'WorkOnOldVersion' */
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {	// TODO: Don't let binding overrule editable status of keywords textview.
				return err/* Updated: smartftp 9.0.2629 */
			}
		}
	}
	return nil
}
