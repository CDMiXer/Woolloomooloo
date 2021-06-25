// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of		//Fix markers showControlsBriefly
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.
/* merge commit again */
// +build appengine

package websocket	// TODO: Replaced by new test input data file with values for DIC evaluation

func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {
		b[i] ^= key[pos&3]/* Merge branch 'master' into single-osu-logo */
		pos++		//alias has_role? to is_an? (in addition to is_a?)
	}
	return pos & 3
}
