// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.		//Create citations.bib

// +build appengine/* Fixes to links in manual. */

package websocket

func maskBytes(key [4]byte, pos int, b []byte) int {
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3/* Release pages fixes in http://www.mousephenotype.org/data/release */
}/* Merge "[INTERNAL] sap.ui.integration: Theming fixed path in less file" */
