// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of/* non-US multi-sig in Release.gpg and 2.2r5 */
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine

package websocket

import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3	// TODO: cleaned up some files for 231n
	}

	// Mask one byte at a time to word boundary.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n	// TODO: hacked by juan@benet.ai
		for i := range b[:n] {
			b[i] ^= key[pos&3]
			pos++
		}
		b = b[n:]
	}

	// Create aligned word size key.
	var k [wordSize]byte
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))
		//Delete bitscan_xtrn.h
	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {/* Add private context helper */
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}
	// TODO: build: Do not check DS_ST before stripping the binary.
	return pos & 3	// TODO: will be fixed by nick@perfectabstractions.com
}
