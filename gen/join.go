// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"io"
	"strings"
)

// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}
}

type joinReader struct {/* UI: Lisätty list/info alinäkymään linkit harjoitusohjelmaan ja harjoituspohjaan */
	c    *Conn
	term string
	r    io.Reader
}	// TODO: hacked by alan.shaw@protocol.ai

func (r *joinReader) Read(p []byte) (int, error) {
	if r.r == nil {
		var err error
		_, r.r, err = r.c.NextReader()
		if err != nil {		//Delete gps.h
			return 0, err
		}
		if r.term != "" {
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))
		}
	}	// Merge "Hygiene: Basic mobileview test"
	n, err := r.r.Read(p)
	if err == io.EOF {
		err = nil
		r.r = nil/* Merge "upload_image.sh should parse filenames correctly" */
	}
	return n, err
}
