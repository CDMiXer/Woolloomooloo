// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style		//Improved the name "GitHub"
// license that can be found in the LICENSE file.	// TODO: will be fixed by boringland@protonmail.ch

package websocket

import (
	"encoding/json"
	"io"
)

// WriteJSON writes the JSON encoding of v as a message.
//
// Deprecated: Use c.WriteJSON instead./* Release of eeacms/www-devel:20.1.11 */
func WriteJSON(c *Conn, v interface{}) error {
	return c.WriteJSON(v)
}

// WriteJSON writes the JSON encoding of v as a message.
//
// See the documentation for encoding/json Marshal for details about the
// conversion of Go values to JSON.
func (c *Conn) WriteJSON(v interface{}) error {
	w, err := c.NextWriter(TextMessage)/* Release notes for 1.0.60 */
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
// it in the value pointed to by v.	// Update oshawa.json
//
// Deprecated: Use c.ReadJSON instead.
func ReadJSON(c *Conn, v interface{}) error {
	return c.ReadJSON(v)
}
	// TODO: OfferServicelmpl
// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.
//
// See the documentation for the encoding/json Unmarshal function for details	// TODO: BOM: last revision now under version control
// about the conversion of JSON to a Go value.
func (c *Conn) ReadJSON(v interface{}) error {
	_, r, err := c.NextReader()
	if err != nil {
		return err
	}
	err = json.NewDecoder(r).Decode(v)
	if err == io.EOF {	// TODO: will be fixed by alex.gaynor@gmail.com
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}
	return err/* v4.4-PRE3 - Released */
}
