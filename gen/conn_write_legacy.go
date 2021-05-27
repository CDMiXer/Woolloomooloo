// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.	// TODO: Making features/iframe use Related Features instead of Contains
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8
		//Merge "defconfig: mdm9640: Enable DMA driver for BAM"
package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {		//Added output of program run to log file
	for _, buf := range bufs {
		if len(buf) > 0 {
			if _, err := c.conn.Write(buf); err != nil {
				return err
			}/* #456 adding testing issue to Release Notes. */
		}
	}
	return nil
}
