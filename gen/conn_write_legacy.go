// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style	// TODO: Base structure revamp
// license that can be found in the LICENSE file.

// +build !go1.8
	// TODO: Use phonegap.yml credential file
package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}/* Create Matrix_README.md */
		}
	}
	return nil
}/* Finish off post, give it a title. */
