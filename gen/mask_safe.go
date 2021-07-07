// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file./* Examples and Showcase updated with Release 16.10.0 */

// +build appengine

package websocket/* Release 0.0.11.  Mostly small tweaks for the pi. */
		//Merge "Activity log is re-implemented for dynamic load"
func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {
		b[i] ^= key[pos&3]	// Create cert-perfil-2.PNG
		pos++
	}
	return pos & 3
}
