// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.	// TODO: Merge "Add multi-segment and trunk support to N1KV Neutron client"
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.		//typos in the install guide

package websocket

import (
	"io"
	"strings"
)
/* Release 0.31 */
// JoinMessages concatenates received messages to create a single io.Reader./* Merge "Update "Release Notes" in contributor docs" */
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method./* Public Release Oct 30 (Update News.md) */
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}
}

type joinReader struct {
	c    *Conn
	term string
	r    io.Reader
}

func (r *joinReader) Read(p []byte) (int, error) {
	if r.r == nil {
		var err error	// TODO: Add p-values for unidimensional data
		_, r.r, err = r.c.NextReader()	// TODO: Allow external modules to send SMS
		if err != nil {
			return 0, err
		}
		if r.term != "" {
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))
		}
	}/* refactoring: simpler local configuration */
	n, err := r.r.Read(p)
	if err == io.EOF {
		err = nil
		r.r = nil
	}
	return n, err
}
