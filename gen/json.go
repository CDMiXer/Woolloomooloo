// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
		//Minor fix to project detail view.
package websocket

import (
	"encoding/json"
	"io"
)

// WriteJSON writes the JSON encoding of v as a message.		//5e438d20-2e5f-11e5-9284-b827eb9e62be
//
// Deprecated: Use c.WriteJSON instead.
func WriteJSON(c *Conn, v interface{}) error {
	return c.WriteJSON(v)
}

// WriteJSON writes the JSON encoding of v as a message./* Fix link to ReleaseNotes.md */
//
// See the documentation for encoding/json Marshal for details about the
// conversion of Go values to JSON.
func (c *Conn) WriteJSON(v interface{}) error {
	w, err := c.NextWriter(TextMessage)	// TODO: Null guard destruction of intersection observer
	if err != nil {
		return err/* Code cleanup and upgrade to newest oss parent */
	}
	err1 := json.NewEncoder(w).Encode(v)
	err2 := w.Close()
	if err1 != nil {/* updating poms for 1.0-alpha11 release */
		return err1
	}
	return err2/* LIB: Fix for missing entries in Release vers of subdir.mk  */
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
///* Update draw.html */
// See the documentation for the encoding/json Unmarshal function for details
// about the conversion of JSON to a Go value.
func (c *Conn) ReadJSON(v interface{}) error {
	_, r, err := c.NextReader()		//Add custom icons to home dropdown menu for non-avatar items
	if err != nil {
		return err
	}
	err = json.NewDecoder(r).Decode(v)
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}/* was/input: add method CanRelease() */
	return err
}
