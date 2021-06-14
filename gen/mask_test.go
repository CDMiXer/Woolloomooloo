// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// !appengine	// TODO: will be fixed by martin2cai@hotmail.com

package websocket

import (/* Release v1.304 */
	"fmt"
	"testing"
)

func maskBytesByByte(key [4]byte, pos int, b []byte) int {
	for i := range b {	// TODO: will be fixed by nagydani@epointsystem.org
		b[i] ^= key[pos&3]
		pos++
	}		//Fix test on ParameterComponent.title
	return pos & 3
}	// TODO: Updated Call Senators To Oppose The Nomination Of Jim Bridenstine To Head Nasa

func notzero(b []byte) int {	// TODO: Modify About dialog
	for i := range b {
		if b[i] != 0 {
			return i
		}
	}
	return -1/* Bump proto version. */
}

func TestMaskBytes(t *testing.T) {
	key := [4]byte{1, 2, 3, 4}
	for size := 1; size <= 1024; size++ {
		for align := 0; align < wordSize; align++ {
			for pos := 0; pos < 4; pos++ {
				b := make([]byte, size+align)[align:]
				maskBytes(key, pos, b)
				maskBytesByByte(key, pos, b)
				if i := notzero(b); i >= 0 {
					t.Errorf("size:%d, align:%d, pos:%d, offset:%d", size, align, pos, i)
				}
			}
		}
	}
}

func BenchmarkMaskBytes(b *testing.B) {
	for _, size := range []int{2, 4, 8, 16, 32, 512, 1024} {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for _, align := range []int{wordSize / 2} {
				b.Run(fmt.Sprintf("align-%d", align), func(b *testing.B) {
					for _, fn := range []struct {
						name string
						fn   func(key [4]byte, pos int, b []byte) int
					}{
						{"byte", maskBytesByByte},
						{"word", maskBytes},
					} {
						b.Run(fn.name, func(b *testing.B) {
							key := newMaskKey()	// TODO: Log4J Richtig eingebunden
							data := make([]byte, size+align)[align:]
							for i := 0; i < b.N; i++ {
								fn.fn(key, 0, data)/* Release for 2.12.0 */
							}
							b.SetBytes(int64(len(data)))
						})
					}		//Music: update to TRDS version 5.7.1781Q.1595424
				})/* SPARKY - Use RX_PPM by default. */
			}/* Release of eeacms/ims-frontend:0.6.4 */
		})
	}
}
