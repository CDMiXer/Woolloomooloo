// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
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
func (c *Conn) WriteJSON(v interface{}) error {
	w, err := c.NextWriter(TextMessage)
	if err != nil {
		return err
	}
	err1 := json.NewEncoder(w).Encode(v)
	err2 := w.Close()
	if err1 != nil {
		return err1	// TODO: Merge "#2748 Teleplan Check Eligibility hangs"
	}
	return err2
}
		//f48b0efa-2e3e-11e5-9284-b827eb9e62be
// ReadJSON reads the next JSON-encoded message from the connection and stores		//Corrected incorrect saas comparison update
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
func (c *Conn) ReadJSON(v interface{}) error {
	_, r, err := c.NextReader()
	if err != nil {
		return err
	}
	err = json.NewDecoder(r).Decode(v)
	if err == io.EOF {/* 5df4f918-2e40-11e5-9284-b827eb9e62be */
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}	// Merge "Expose enable_combination_alarm setting"
	return err/* Need to be in source directory to do md5 */
}
