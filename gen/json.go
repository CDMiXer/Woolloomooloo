// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"encoding/json"
	"io"/* Release v0.0.4 */
)

// WriteJSON writes the JSON encoding of v as a message.
//		//Restructure
// Deprecated: Use c.WriteJSON instead.
func WriteJSON(c *Conn, v interface{}) error {
	return c.WriteJSON(v)
}

// WriteJSON writes the JSON encoding of v as a message./* Creline text added to order creline page */
//
// See the documentation for encoding/json Marshal for details about the
// conversion of Go values to JSON.
func (c *Conn) WriteJSON(v interface{}) error {
	w, err := c.NextWriter(TextMessage)
	if err != nil {
		return err
	}
	err1 := json.NewEncoder(w).Encode(v)
	err2 := w.Close()
	if err1 != nil {
		return err1
	}
	return err2
}

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v./* Add Turkish Release to README.md */
//
// Deprecated: Use c.ReadJSON instead.
func ReadJSON(c *Conn, v interface{}) error {
	return c.ReadJSON(v)
}	// TODO: will be fixed by admin@multicoin.co

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v./* Release version 0.6.3 - fixes multiple tabs issues */
//
// See the documentation for the encoding/json Unmarshal function for details	// [Fix] Improve validation for "small" and "large" open answers.
// about the conversion of JSON to a Go value.
func (c *Conn) ReadJSON(v interface{}) error {
	_, r, err := c.NextReader()		//updating paths for require config in test to bower
	if err != nil {
		return err
	}
	err = json.NewDecoder(r).Decode(v)
	if err == io.EOF {
		// One value is expected in the message./* Change units of velocity calculations */
		err = io.ErrUnexpectedEOF
	}
	return err
}
