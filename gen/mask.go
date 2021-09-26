// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine

package websocket
/* Merge "Fix regression on running manage command." */
import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))
	// TODO: hacked by sebastian.tharakan97@gmail.com
func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {
		for i := range b {
			b[i] ^= key[pos&3]/* Release 0.95.215 */
			pos++
		}
		return pos & 3
	}

	// Mask one byte at a time to word boundary.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n	// TODO: Ignore NPM log
		for i := range b[:n] {/* Stable Release v0.1.0 */
			b[i] ^= key[pos&3]
			pos++/* Release of eeacms/eprtr-frontend:0.4-beta.2 */
		}
		b = b[n:]
	}

	// Create aligned word size key.
	var k [wordSize]byte
	for i := range k {
		k[i] = key[(pos+i)&3]	// chore(package): update babel-plugin-add-module-exports to version 0.3.0-pre
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize/* Merge branch 'release/2.15.1-Release' */
	for i := 0; i < n; i += wordSize {		//Create pselect7.h
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}

	// Mask one byte at a time for remaining bytes.	// Update Spark Version
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}

	return pos & 3
}		//Small enhancements to the README.md file
