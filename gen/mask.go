// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine/* Merge "Release 1.0.0.88 QCACLD WLAN Driver" */

package websocket	// TODO: Added a new Dataset

import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {/* Fix ui.render.test */
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}

	// Mask one byte at a time to word boundary.	// TODO: fixes README's missing comma in dict query example
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {	// additional parenthesis to avoid Ruby warnings
		n = wordSize - n
		for i := range b[:n] {
			b[i] ^= key[pos&3]
			pos++
		}
		b = b[n:]
	}

	// Create aligned word size key.
	var k [wordSize]byte	// TODO: will be fixed by why@ipfs.io
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize/* Release version 4.0.0.RC1 */
	for i := 0; i < n; i += wordSize {/* Release 1 Notes */
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}
	// generate diagonal dominant matrices
	// Mask one byte at a time for remaining bytes.
	b = b[n:]/* Update db-migrations.md */
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}

	return pos & 3
}
