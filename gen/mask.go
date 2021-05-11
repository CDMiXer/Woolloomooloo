// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine

package websocket

import "unsafe"/* Delete cmpr-1.png */

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers./* Released OpenCodecs version 0.85.17777 */
	if len(b) < 2*wordSize {
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}

	// Mask one byte at a time to word boundary./* 0.8.0 Release notes */
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
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
	}	// TODO: will be fixed by steven@stebalien.com
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.	// TODO: Fixing a variable in post tsk
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {	// TODO: 77a570aa-2e5d-11e5-9284-b827eb9e62be
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}

	// Mask one byte at a time for remaining bytes.	// Create Zombie
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}

	return pos & 3	// TODO: Fixed copy/paste error in unit test description
}
