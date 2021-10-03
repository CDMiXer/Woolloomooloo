// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the	// TODO: will be fixed by 13860583249@yeah.net
// LICENSE file.

// +build appengine
		//fixed run environment to windows
package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}/* Rename Heit el-Ghurab to Heit el-Ghurab.html */
	return pos & 3
}
