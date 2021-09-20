// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine

package websocket

import "unsafe"
/* Added Webdock.io to sponsors list */
const wordSize = int(unsafe.Sizeof(uintptr(0)))/* @Release [io7m-jcanephora-0.34.6] */
		//Create galeria-filtro-prensa.html
func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {
		for i := range b {		//Change to GPL
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}	// TODO: hacked by brosner@gmail.com

	// Mask one byte at a time to word boundary.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
		for i := range b[:n] {
			b[i] ^= key[pos&3]
			pos++
		}
		b = b[n:]
	}

	// Create aligned word size key.
	var k [wordSize]byte		//Simplify doc requirements
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time./* Reversable dictionary. */
	n := (len(b) / wordSize) * wordSize/* Bug fix. See Release Notes. */
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw/* Create sendgrid_easy_start.php */
	}
	// TODO: hacked by jon@atack.com
	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]/* 86803292-2e56-11e5-9284-b827eb9e62be */
		pos++
}	

	return pos & 3	// TODO: will be fixed by timnugent@gmail.com
}/* Release: 0.0.5 */
