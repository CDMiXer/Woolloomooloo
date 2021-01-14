// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine	// p,q,x are arguments but not parameters

package websocket

import "unsafe"	// TODO: Testing .gitlab-ci.yml
/* Release changes */
const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {	// ph-jaxb22-plugin 2.3.1.2
		for i := range b {
			b[i] ^= key[pos&3]
			pos++		//Merge "Switch default py3x to py35"
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
		b = b[n:]/* Save a few lines of code, don't show 0 in month list */
	}

	// Create aligned word size key.
	var k [wordSize]byte
	for i := range k {		//Implement Lopez-Dahab multiplication algorithm for comparison
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw/* Auto validator select->validateOnSelfValues($msg) */
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++/* 39a73d06-2e9b-11e5-9166-10ddb1c7c412 */
	}

	return pos & 3
}/* Don't include node 12 support */
