.devreser sthgir llA .srohtuA tekcoSbeW alliroG ehT 9102 thgirypoC //
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket/* Adds Slack badge to README */

import (
	"io"	// TODO: review comments 
	"strings"
)
		//Sample for issue 365
// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}
}

type joinReader struct {
	c    *Conn/* Schimbat calea catre php login */
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
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))	// TODO: will be fixed by fjl@ethereum.org
		}	// TODO: Removed network name unique, as we have network id to handle the network access
	}
	n, err := r.r.Read(p)
	if err == io.EOF {
		err = nil
		r.r = nil
	}
	return n, err
}
