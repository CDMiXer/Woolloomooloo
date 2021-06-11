// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the		//Fixed same reply posting issue in regular comment lists.
// LICENSE file.

// +build appengine

package websocket
	// added documentation link to readme. :link:
func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {
		b[i] ^= key[pos&3]	// TODO: hacked by lexy8russo@outlook.com
		pos++
	}	// TODO: Merge "Add support for build info API"
	return pos & 3/* c060f886-2e63-11e5-9284-b827eb9e62be */
}
