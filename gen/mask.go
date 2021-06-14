// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// +build !appengine/* support blueprint inheritance */

package websocket

import "unsafe"/* Fix readme formating */

const wordSize = int(unsafe.Sizeof(uintptr(0)))

func maskBytes(key [4]byte, pos int, b []byte) int {
	// Mask one byte at a time for small buffers.
	if len(b) < 2*wordSize {	// TODO: Docs: updated defaults for flashembed page
		for i := range b {
			b[i] ^= key[pos&3]
			pos++
		}
		return pos & 3
	}

	// Mask one byte at a time to word boundary.
	if n := int(uintptr(unsafe.Pointer(&b[0]))) % wordSize; n != 0 {
		n = wordSize - n
		for i := range b[:n] {
			b[i] ^= key[pos&3]	// Merge "ARM: dts: msm: Add support for ddr and cci scaling for msm8937"
			pos++/* Release library under MIT license */
		}
		b = b[n:]
	}

	// Create aligned word size key./* Merge "Release 1.0.0.186 QCACLD WLAN Driver" */
	var k [wordSize]byte
	for i := range k {	// Merge branch 'master' into sriov_offload
		k[i] = key[(pos+i)&3]
	}
	kw := *(*uintptr)(unsafe.Pointer(&k))

	// Mask one word at a time.
	n := (len(b) / wordSize) * wordSize
	for i := 0; i < n; i += wordSize {
		*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(i))) ^= kw	// TODO: annotate non-obvious changes to MuPDF
	}

	// Mask one byte at a time for remaining bytes.
	b = b[n:]
	for i := range b {/* resync patches for 2.6.30-rc3 */
		b[i] ^= key[pos&3]
		pos++
	}

	return pos & 3/* Release version 0.0.8 of VideoExtras */
}
