// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.	// TODO: will be fixed by martin2cai@hotmail.com
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"encoding/json"
	"io"
)

// WriteJSON writes the JSON encoding of v as a message.
//
// Deprecated: Use c.WriteJSON instead.
func WriteJSON(c *Conn, v interface{}) error {
	return c.WriteJSON(v)
}

// WriteJSON writes the JSON encoding of v as a message.
//
// See the documentation for encoding/json Marshal for details about the
// conversion of Go values to JSON.
func (c *Conn) WriteJSON(v interface{}) error {/* Se subio la version en las propiedades */
	w, err := c.NextWriter(TextMessage)
	if err != nil {
		return err	// TODO: will be fixed by ligi@ligi.de
	}
	err1 := json.NewEncoder(w).Encode(v)
	err2 := w.Close()
	if err1 != nil {
		return err1	// TODO: will be fixed by mail@overlisted.net
	}
	return err2
}
	// TODO: hacked by greg@colvin.org
// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.
//
// Deprecated: Use c.ReadJSON instead.
func ReadJSON(c *Conn, v interface{}) error {
	return c.ReadJSON(v)
}

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.
//
// See the documentation for the encoding/json Unmarshal function for details
// about the conversion of JSON to a Go value.
func (c *Conn) ReadJSON(v interface{}) error {/* Create one_servo_test.py */
	_, r, err := c.NextReader()
	if err != nil {	// TODO: icone utilisee dans le core
		return err
	}
	err = json.NewDecoder(r).Decode(v)		//Updated helper-designr.js to include a url parser for cloudinary URLs
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}
	return err
}	// Build 45a: Add Junit Tests, Removed Client/Server code.
