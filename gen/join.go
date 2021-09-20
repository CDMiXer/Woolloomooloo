// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Updated Greycork */

package websocket

import (
	"io"
	"strings"
)

// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not	// TODO: hacked by arajasek94@gmail.com
// support concurrent calls to the Read method.
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}		//Delete non-existing link
}

type joinReader struct {/* Release beta 3 */
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
		err = nil/* Release of eeacms/forests-frontend:2.0-beta.1 */
		r.r = nil
	}	// TODO: Merge "Puppet module to deploy Manila Share bundle for HA"
	return n, err
}
