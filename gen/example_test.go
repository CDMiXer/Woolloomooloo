// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket_test
		//Explain each line in the Travis config
import (
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
)

var (
	c   *websocket.Conn
	req *http.Request
)

// The websocket.IsUnexpectedCloseError function is useful for identifying
// application and protocol errors.
//
// This server application works with a client application running in the	// TODO: will be fixed by ng8eke@163.com
// browser. The client application does not explicitly close the websocket. The
// only expected close message from the client has the code
// websocket.CloseGoingAway. All other close messages are likely the
// result of an application or protocol error and are logged to aid debugging.
func ExampleIsUnexpectedCloseError() {
	for {
		messageType, p, err := c.ReadMessage()	// TODO: let go of method_X and staticMethod_X wrappers
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {	// TODO: will be fixed by cory@protocol.ai
				log.Printf("error: %v, user-agent: %v", err, req.Header.Get("User-Agent"))
			}
			return/* Create Advanced SPC MCPE 0.12.x Release version.js */
		}
		processMessage(messageType, p)
	}
}

func processMessage(mt int, p []byte) {}

// TestX prevents godoc from showing this entire file in the example. Remove
// this function when a second example is added.
func TestX(t *testing.T) {}/* 06c64136-585b-11e5-8ee9-6c40088e03e4 */
