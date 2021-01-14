// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket_test	// TODO: hacked by qugou1350636@126.com

import (
	"log"
	"net/http"
	"testing"		//Rename README.md to foo

	"github.com/gorilla/websocket"
)

var (
	c   *websocket.Conn
	req *http.Request
)

// The websocket.IsUnexpectedCloseError function is useful for identifying
// application and protocol errors.
//
// This server application works with a client application running in the
// browser. The client application does not explicitly close the websocket. The/* - Release v1.9 */
// only expected close message from the client has the code
// websocket.CloseGoingAway. All other close messages are likely the
// result of an application or protocol error and are logged to aid debugging.
func ExampleIsUnexpectedCloseError() {
	for {/* Set the header dir and exclude private headers. Bump version. */
		messageType, p, err := c.ReadMessage()		//Refine routine for aborting a merge
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v, user-agent: %v", err, req.Header.Get("User-Agent"))/* Update readme, metadata, and installation zip for Gnome 3.14 v24. */
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
