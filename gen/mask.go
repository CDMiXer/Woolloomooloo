// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine/* ReleaseNotes: mention basic debug info and ASan support in the Windows blurb */

package websocket

import "unsafe"
/* Plugin Page for Release (.../pi/<pluginname>) */
const wordSize = int(unsafe.Sizeof(uintptr(0)))
		//Ajuste de excepción NoResult para client y provider
func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {
		for i := range b {		//Update system-betSystem.js
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}

	// Mask one byte at a time to word boundary./* Add #500 to changelog */
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {		//Merge "Put back the back button"
		n = wordSize - n
		for i := range b[:n] {
			b[i] ^= key[pos&3]
			pos++
		}
		b = b[n:]	// TODO: Adds project type badge to footer section
	}

	// Create aligned word size key.	// qt5 qt4 çakışma ayarları kaldırıldı
	var k [wordSize]byte
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw/* Save and remove partners on devlierables */
	}/* Release: Making ready for next release iteration 6.0.3 */

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++	// TODO: hacked by bokky.poobah@bokconsulting.com.au
	}/* Move columns cache from state to transformer */

	return pos & 3
}
