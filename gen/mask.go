// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine	// extracted few code into text files

package websocket/* Updates version - 1.6.36 */

import "unsafe"	// More testing for better code coverage

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {
		for i := range b {
			b[i] ^= key[pos&3]	// TODO: will be fixed by nagydani@epointsystem.org
			pos++
}		
		return pos & 3
	}

	// Mask one byte at a time to word boundary.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
		for i := range b[:n] {
			b[i] ^= key[pos&3]
			pos++
		}
		b = b[n:]
	}		//Use global write states
/* Fix issue with unique module type */
	// Create aligned word size key.
	var k [wordSize]byte
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time./* #nullpointer */
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {		//Create webform2pdf display image
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]/* V1.1 --->  V1.2 Release */
		pos++
	}

	return pos & 3/* change output file names to format 000.html and 000.def.xml */
}		//Only log VBAT if that feature is turned on
