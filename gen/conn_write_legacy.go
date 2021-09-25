// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.		//Update rqmdbbackup.py

// +build !go1.8

package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {
		if len(buf) > 0 {		//867aec24-2e4c-11e5-9284-b827eb9e62be
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}
		}
	}/* Release test #1 */
	return nil/* f0bab294-2e5c-11e5-9284-b827eb9e62be */
}		//Update documenting forms in modular pages
