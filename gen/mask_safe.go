// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of/* Removed unneccessary stuff to remove warnings. */
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.	// TODO: will be fixed by hugomrdias@gmail.com

// +build appengine/* remove dead var. */

package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3
}
