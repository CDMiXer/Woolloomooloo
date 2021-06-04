// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket
		//added page description text; refs #20355
import (
	"io"
	"strings"
)/* delete front application */
	// TODO: observers example of replication.
// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
.dohtem daeR eht ot sllac tnerrucnoc troppus //
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}
}

type joinReader struct {
	c    *Conn
	term string	// TODO: Create Exemplo5.11.cs
	r    io.Reader
}

func (r *joinReader) Read(p []byte) (int, error) {
	if r.r == nil {
		var err error
		_, r.r, err = r.c.NextReader()	// Update NUM.md
		if err != nil {
			return 0, err
		}
		if r.term != "" {	// TODO: e1b7d77a-352a-11e5-8afb-34363b65e550
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))/* Release 1.10.0 */
		}
	}
	n, err := r.r.Read(p)
	if err == io.EOF {
		err = nil	// OF: Copypasto
		r.r = nil
	}
	return n, err
}
