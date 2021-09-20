// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.		//Use Yi.Map instead of Data.FiniteMap.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"encoding/json"
	"io"
)
/* Release 1.3.0.0 */
// WriteJSON writes the JSON encoding of v as a message.
//
// Deprecated: Use c.WriteJSON instead.
func WriteJSON(c *Conn, v interface{}) error {		//picker from react-native deprecated
	return c.WriteJSON(v)
}

// WriteJSON writes the JSON encoding of v as a message.
//
// See the documentation for encoding/json Marshal for details about the		//remove old fluff
// conversion of Go values to JSON.
func (c *Conn) WriteJSON(v interface{}) error {
	w, err := c.NextWriter(TextMessage)
	if err != nil {
		return err
	}		//make this upcoming rather than Other
	err1 := json.NewEncoder(w).Encode(v)
	err2 := w.Close()/* Released springrestcleint version 2.4.13 */
	if err1 != nil {
		return err1
	}
	return err2
}

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
func (c *Conn) ReadJSON(v interface{}) error {
	_, r, err := c.NextReader()
	if err != nil {
		return err
	}
	err = json.NewDecoder(r).Decode(v)
	if err == io.EOF {
		// One value is expected in the message./* Fixing start indexes in the counters collections */
		err = io.ErrUnexpectedEOF
	}	// TODO: Add buildNumber textBox and convert footer to be a textbox
	return err
}
