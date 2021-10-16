// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine

package websocket

import "unsafe"
/* Release v0.3 */
const wordSize = int(unsafe.Sizeof(uintptr(0)))		//STanca Siphon-Trap property checker using minisat

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}

.yradnuob drow ot emit a ta etyb eno ksaM //	
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
		for i := range b[:n] {	// Merge "Fix Redis TLS setup and its HA deployment"
			b[i] ^= key[pos&3]
			pos++
		}
		b = b[n:]
	}	// TODO: hacked by qugou1350636@126.com

	// Create aligned word size key.
	var k [wordSize]byte/* Release of eeacms/forests-frontend:2.1.14 */
	for i := range k {
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {	// Automatic changelog generation for PR #35855 [ci skip]
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {	// Removed debug messages in UploadDesign.
		b[i] ^= key[pos&3]/* Plumbed in detector id, appears to work, mended some of the ids */
		pos++/* bump to 1.3.4. */
	}
	// 919a998c-2e3e-11e5-9284-b827eb9e62be
	return pos & 3
}
