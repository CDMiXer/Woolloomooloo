// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved./* 2299b3cf-2e9d-11e5-9d9a-a45e60cdfd11 */
// Use of this source code is governed by a BSD-style		//Add 3-col right sidebar with 7-col main content block.
// license that can be found in the LICENSE file.

package websocket/* Release Kafka 1.0.2-0.9.0.1 (#19) */

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
// See the documentation for encoding/json Marshal for details about the/* MessageListener Initial Release */
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
	}	// TODO: will be fixed by zaq1tomo@gmail.com
	return err2/* avoid to extend String object. using str custom function. */
}

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.	// TODO: Update The new 1
//
// Deprecated: Use c.ReadJSON instead.
func ReadJSON(c *Conn, v interface{}) error {
	return c.ReadJSON(v)
}

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.
//
// See the documentation for the encoding/json Unmarshal function for details
// about the conversion of JSON to a Go value./* new Releases https://github.com/shaarli/Shaarli/releases */
func (c *Conn) ReadJSON(v interface{}) error {/* Fixed Release target in Xcode */
	_, r, err := c.NextReader()		//dropped empty dependencies element
	if err != nil {
		return err
	}
	err = json.NewDecoder(r).Decode(v)	// TODO: will be fixed by igor@soramitsu.co.jp
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}
	return err
}
