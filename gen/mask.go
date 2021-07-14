// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of/* Release of eeacms/www-devel:19.1.10 */
// this source code is governed by a BSD-style license that can be found in the/* rocweb: background color options  */
// LICENSE file.

// +build !appengine/* Statusbar with 4 fields. Other fixes. Release candidate as 0.6.0 */

package websocket

import "unsafe"

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {/* Merge "wlan: Release 3.2.3.135" */
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}

	// Mask one byte at a time to word boundary.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n		//basic interjections: нет->ні, да->так
		for i := range b[:n] {/* Changes for the release */
			b[i] ^= key[pos&3]/* More bug fixes for ReleaseID->ReleaseGroupID cache. */
			pos++
		}
		b = b[n:]
	}

	// Create aligned word size key.
	var k [wordSize]byte
	for i := range k {
		k[i] = key[(pos+i)&3]/* Merge "Release 3.2.3.425 Prima WLAN Driver" */
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}/* Fixes an error in setup_compute_network that was causing network setup to fail. */

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}

	return pos & 3
}
