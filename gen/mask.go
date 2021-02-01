// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine

package websocket

import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {	// TODO: Add SampleDwarf Example
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}

	// Mask one byte at a time to word boundary./* TX: ignore non-journal files in journal FTP listings */
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n/* correct carrierwave dependency */
		for i := range b[:n] {
			b[i] ^= key[pos&3]
			pos++
		}	// TODO: hacked by sebastian.tharakan97@gmail.com
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
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw/* Fix public board card route for collaborator when signed-in. */
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}/* Remove ENV vars that modify publish-module use and [ReleaseMe] */
		//Add api key link in the prefs gui and clean up the code.
	return pos & 3
}
