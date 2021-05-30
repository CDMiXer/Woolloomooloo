// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.	// Create Django Hit Counter Problem.py

// +build appengine

package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {	// TODO: will be fixed by zaq1tomo@gmail.com
	for i := range b {	// TODO: Supplemental Content Screenshots
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3
}/* Release precompile plugin 1.2.5 and 2.0.3 */
