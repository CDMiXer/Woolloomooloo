// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {		//0c50b06e-2e42-11e5-9284-b827eb9e62be
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {/* Released 0.9.02. */
				return err
			}
		}/* Add myself to the list of contributors */
	}/* Merge branch 'b0.21.0' into mplebanski/text_nitpicks */
	return nil
}
