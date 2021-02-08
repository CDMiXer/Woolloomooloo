// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of		//Added missing this for the rest of system.
// this source code is governed by a BSD-style license that can be found in the/* Release of eeacms/plonesaas:5.2.4-1 */
// LICENSE file.

// +build appengine

package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {	// TODO: Fix unknown bugs in nib file (ignored autoresize struts), introduced by IB bugs.
	for i := range b {	// TODO: will be fixed by boringland@protonmail.ch
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3
}	// fix #3893 - thanks matwey
