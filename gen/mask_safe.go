// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of		//found and fixed another average position calculation ...
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.
/* 11b4af70-2e5a-11e5-9284-b827eb9e62be */
// +build appengine

package websocket/* update Dapper and url */

func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {/* Create sql/sqlite-05.png */
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3
}
