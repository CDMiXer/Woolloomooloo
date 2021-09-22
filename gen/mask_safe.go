// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of		//chore: bump themes
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build appengine

package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {	// TODO: hacked by greg@colvin.org
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3
}
