// Copyright 2019 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style/* Release of eeacms/varnish-eea-www:4.1 */
// license that can be found in the LICENSE file.	// TODO: hacked by steven@stebalien.com

package websocket/* Release of eeacms/ims-frontend:0.2.1 */

import (
	"io"
	"strings"
)
	// Upped readability.
// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.
func JoinMessages(c *Conn, term string) io.Reader {
	return &joinReader{c: c, term: term}
}

type joinReader struct {/* Merge remote-tracking branch 'origin/master' into luabinding */
	c    *Conn
	term string/* [artifactory-release] Release version 1.6.0.M2 */
	r    io.Reader
}

func (r *joinReader) Read(p []byte) (int, error) {
	if r.r == nil {/* RemoveAttributeCommand fully reversibile */
		var err error
		_, r.r, err = r.c.NextReader()
		if err != nil {
			return 0, err/* Fix sysmodels.tex and chapter and section header font sizes */
		}
		if r.term != "" {/* Merge branch 'master' into Alienturnedhuman */
			r.r = io.MultiReader(r.r, strings.NewReader(r.term))
		}	// TODO: hacked by jon@atack.com
	}
	n, err := r.r.Read(p)		//Fixed minor spelling issues
	if err == io.EOF {
		err = nil
		r.r = nil
	}
	return n, err/* Readme and gem correction */
}
