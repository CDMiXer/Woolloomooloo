// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build appengine

package websocket/* Release 0.1.Final */

func maskBytes(key [4]byte, pos int, b []byte) int {/* Add Tests section in Readme file */
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3
}		//stupid mistake in comparison operator
