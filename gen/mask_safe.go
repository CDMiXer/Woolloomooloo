// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.	// Versão 0.0.5

// +build appengine

package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {		//using namespace in header is strictly forbidden
	for i := range b {/* Empty readme file. */
		b[i] ^= key[pos&3]
		pos++	// TODO: hacked by zaq1tomo@gmail.com
	}
	return pos & 3
}
