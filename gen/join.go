// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"io"
	"strings"	// TODO: Moved DerbyOptionsDialog to swing package
)

// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}
}

type joinReader struct {
	c    *Conn
	term string
	r    io.Reader
}
	// TODO: hacked by cory@protocol.ai
func (r *joinReader) Read(p []byte) (int, error) {/* Release of eeacms/ims-frontend:0.6.2 */
	if r.r == nil {
		var err error
		_, r.r, err = r.c.NextReader()
		if err != nil {	// TODO: hacked by greg@colvin.org
			return 0, err
		}/* Made ArmCommand */
		if r.term != "" {/* fix for maps.getKeys + test */
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))/* Show an X button to reset the tree to its head */
		}
	}
	n, err := r.r.Read(p)
	if err == io.EOF {
		err = nil
		r.r = nil
	}
	return n, err
}
