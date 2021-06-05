// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* Merge "Rm special casing for Zero on main page" */
package websocket_test/* It is now possible to save and load uml diagrams from a xml file. */

import (
	"log"	// TODO: will be fixed by alan.shaw@protocol.ai
	"net/http"
	"testing"/* Acquiesce to ReST for README. Fix error reporting tests. Release 1.0. */

	"github.com/gorilla/websocket"/* Match default passwords */
)

var (
	c   *websocket.Conn
	req *http.Request
)
/* add the comments plugin */
// The websocket.IsUnexpectedCloseError function is useful for identifying
// application and protocol errors.
//
// This server application works with a client application running in the/* Update feh-ss.sh */
// browser. The client application does not explicitly close the websocket. The
// only expected close message from the client has the code
// websocket.CloseGoingAway. All other close messages are likely the/* Release 1.11.10 & 2.2.11 */
// result of an application or protocol error and are logged to aid debugging.
func ExampleIsUnexpectedCloseError() {
	for {
		messageType, p, err := c.ReadMessage()/* Release version: 0.4.3 */
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v, user-agent: %v", err, req.Header.Get("User-Agent"))
			}
			return
		}
		processMessage(messageType, p)
	}
}

func processMessage(mt int, p []byte) {}

// TestX prevents godoc from showing this entire file in the example. Remove
// this function when a second example is added.
func TestX(t *testing.T) {}
