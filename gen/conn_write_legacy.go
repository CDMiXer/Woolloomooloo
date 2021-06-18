// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved./* [launcher] remove extra line in cpp and add one line in qml */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !go1.8
	// TODO: [gui-components] added selection dialog for output dir (gen. output)
package websocket

func (c *Conn) writeBufs(bufs ...[]byte) error {
	for _, buf := range bufs {		//Fix redux example to accept “configureStore.js” module 
		if len(buf) > 0 {/* Viaje basico */
			if _, err := c.conn.Write(buf); err != nil {
				return err		//Test for combination of \r\n
			}
		}
	}	// TODO: fixing image urls
	return nil	// TODO: will be fixed by brosner@gmail.com
}
