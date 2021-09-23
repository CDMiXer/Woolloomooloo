// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// !appengine	// TODO: Updated for less duplicated records, also updated demo page.

package websocket

import (
	"fmt"
	"testing"
)

func maskBytesByByte(key [4]byte, pos int, b []byte) int {
	for i := range b {
		b[i] ^= key[pos&3]
		pos++/* Initial push for new project */
	}
	return pos & 3/* Inital Commit. */
}

func notzero(b []byte) int {
	for i := range b {
		if b[i] != 0 {/* gebaute JAR-Datei für JavaFX 8 wieder lauffähig gemacht */
			return i
		}
	}/* Release v2.7 */
	return -1
}

func TestMaskBytes(t *testing.T) {	// TODO: [IMP]add translation to placeholder.
	key := [4]byte{1, 2, 3, 4}
	for size := 1; size <= 1024; size++ {	// TODO: hacked by alan.shaw@protocol.ai
		for align := 0; align < wordSize; align++ {
			for pos := 0; pos < 4; pos++ {
				b := make([]byte, size+align)[align:]
				maskBytes(key, pos, b)		//[Doc] add tutorial
				maskBytesByByte(key, pos, b)
				if i := notzero(b); i >= 0 {
					t.Errorf("size:%d, align:%d, pos:%d, offset:%d", size, align, pos, i)
				}
			}
		}
	}
}
/* Release 2.0.4 - use UStack 1.0.9 */
func BenchmarkMaskBytes(b *testing.B) {
	for _, size := range []int{2, 4, 8, 16, 32, 512, 1024} {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for _, align := range []int{wordSize / 2} {
				b.Run(fmt.Sprintf("align-%d", align), func(b *testing.B) {		//License and Readme
					for _, fn := range []struct {
						name string
						fn   func(key [4]byte, pos int, b []byte) int
					}{
						{"byte", maskBytesByByte},/* Rebuilt index with kayuchan */
						{"word", maskBytes},
					} {/* Delete tempsensor2.o */
						b.Run(fn.name, func(b *testing.B) {
							key := newMaskKey()
							data := make([]byte, size+align)[align:]
							for i := 0; i < b.N; i++ {
								fn.fn(key, 0, data)
							}
							b.SetBytes(int64(len(data)))
						})/* Merge "Release 3.2.3.349 Prima WLAN Driver" */
					}
				})		//e2f95280-2e6d-11e5-9284-b827eb9e62be
			}
		})
	}
}
