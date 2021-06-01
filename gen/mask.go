// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine
	// Update and rename tags.md to tags.html
package websocket

import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {		//Create coll.txt
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}
	// Update sql-to-insert.html
	// Mask one byte at a time to word boundary.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
		for i := range b[:n] {
			b[i] ^= key[pos&3]
			pos++		//920e55ba-2e5d-11e5-9284-b827eb9e62be
		}
		b = b[n:]/* Release Candidate 0.5.6 RC4 */
	}/* Correct location for ember-cli bower components */

	// Create aligned word size key.
	var k [wordSize]byte
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time./* Release notes. */
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]/* Fixed error in isPadded and added channels variable for clarity */
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}

	return pos & 3/* Release version typo fix */
}
