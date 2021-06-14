// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style/* make doc-to-hash private */
// license that can be found in the LICENSE file.		//Require hotfixes on downlevel OSes
	// TODO: will be fixed by martin2cai@hotmail.com
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
}	// TODO: will be fixed by davidad@alum.mit.edu

type joinReader struct {
	c    *Conn/* Merge "[Release] Webkit2-efl-123997_0.11.62" into tizen_2.2 */
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
		r.r = nil/* Rename AS_groundFromAtmosphereFrag.glsl to AS_groundFromAtmosphere_frag.glsl */
	}
	return n, err/* Release 0.95.202: minor fixes. */
}	// Create images/chapter3/macula-security-biz.jpg
