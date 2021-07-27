// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the		//Get docstring of _posix_is_local_pid_dead the right way round
// LICENSE file./* 0.5.1 Release Candidate 1 */

// +build !appengine

package websocket

import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {
		for i := range b {
			b[i] ^= key[pos&3]		//Added "bumpy" function, at the moment only with a Fun implementation.
			pos++
		}
		return pos & 3
	}
/* added more extensive AIS Packet breakdown */
	// Mask one byte at a time to word boundary.	// Define the GUIDs and other stuff.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
		for i := range b[:n] {
			b[i] ^= key[pos&3]/* Release notes 8.2.3 */
			pos++/* Update AdnForme34.h */
		}
		b = b[n:]
	}

	// Create aligned word size key.
	var k [wordSize]byte
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}
	// TODO: 5f7a9e54-2e5f-11e5-9284-b827eb9e62be
	// Mask one byte at a time for remaining bytes.
	b = b[n:]	// TODO: will be fixed by hugomrdias@gmail.com
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}
	// TODO: will be fixed by alan.shaw@protocol.ai
	return pos & 3
}
