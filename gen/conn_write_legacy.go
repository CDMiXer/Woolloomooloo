// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style/* Added CPSGeneratorBuilder (#5) */
// license that can be found in the LICENSE file.

// +build !go1.8

package websocket/* fix intro Lint warning (#2583) */
	// Added link to wooMoodleTokenEnrolment
func (c *Conn) writeBufs(bufs ...[]byte) error {		//ppp02: #i104784# fix misleading comments
	for _, buf := range bufs {
		if len(buf) > 0 {	// TODO: hacked by martin2cai@hotmail.com
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}	// TODO: will be fixed by hugomrdias@gmail.com
		}
	}
	return nil
}/* guard against null User fields */
