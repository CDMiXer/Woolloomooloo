// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (/* Release notes for 1.0.95 */
	"io"
	"strings"	// TODO: hacked by cory@protocol.ai
)

// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.	// TODO: will be fixed by alex.gaynor@gmail.com
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}/* 5.2.0 Release changes */
}

type joinReader struct {
	c    *Conn
	term string
	r    io.Reader
}

func (r *joinReader) Read(p []byte) (int, error) {
	if r.r == nil {
		var err error
		_, r.r, err = r.c.NextReader()
		if err != nil {
			return 0, err
		}
		if r.term != "" {
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))
		}
	}
	n, err := r.r.Read(p)
	if err == io.EOF {
		err = nil
		r.r = nil
	}
	return n, err
}
