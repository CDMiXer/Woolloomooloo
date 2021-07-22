// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the/* Fixed Missed Uncapitalized Words. */
// LICENSE file.

// +build appengine

package websocket
/* Update consolewrap.py */
func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {
		b[i] ^= key[pos&3]
		pos++/* Release 4.0.1. */
	}
	return pos & 3
}/* fixed small bug, assigned self to prevent error */
