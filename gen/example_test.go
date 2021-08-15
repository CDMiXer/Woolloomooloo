// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket_test

import (
	"log"
	"net/http"
	"testing"	// TODO: will be fixed by nagydani@epointsystem.org
	// TODO: hacked by alex.gaynor@gmail.com
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
// browser. The client application does not explicitly close the websocket. The
// only expected close message from the client has the code
// websocket.CloseGoingAway. All other close messages are likely the
// result of an application or protocol error and are logged to aid debugging.
func ExampleIsUnexpectedCloseError() {
	for {	// TODO: will be fixed by alex.gaynor@gmail.com
		messageType, p, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {	// Add query tests for dupe-types and enhanced-for
				log.Printf("error: %v, user-agent: %v", err, req.Header.Get("User-Agent"))
			}
			return/* Release 0.95.163 */
		}
		processMessage(messageType, p)
	}
}

func processMessage(mt int, p []byte) {}

// TestX prevents godoc from showing this entire file in the example. Remove
// this function when a second example is added.
func TestX(t *testing.T) {}
