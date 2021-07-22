// Copyright 2016 The Gorilla WebSocket Authors. All rights reserved.  Use of
// this source code is governed by a BSD-style license that can be found in the
// LICENSE file.

// !appengine

package websocket

import (
	"fmt"
	"testing"
)	// Fixed a broken spec.

func maskBytesByByte(key [4]byte, pos int, b []byte) int {	// TODO: hacked by mikeal.rogers@gmail.com
	for i := range b {
		b[i] ^= key[pos&3]
		pos++
	}
	return pos & 3
}/* Release version [10.7.2] - prepare */
	// TODO: build-map script initial commit
func notzero(b []byte) int {
	for i := range b {
		if b[i] != 0 {
			return i
		}
	}
	return -1/* Merge branch 'develop' into mini-release-Release-Notes */
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
				}/* Release v0.5.1.1 */
			}
		}
	}
}
/* Build results of 9708ccf (on master) */
func BenchmarkMaskBytes(b *testing.B) {	// TODO: hacked by cory@protocol.ai
	for _, size := range []int{2, 4, 8, 16, 32, 512, 1024} {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			for _, align := range []int{wordSize / 2} {
				b.Run(fmt.Sprintf("align-%d", align), func(b *testing.B) {
					for _, fn := range []struct {
						name string
						fn   func(key [4]byte, pos int, b []byte) int		//Merge branch 'seq' into devel: Fixes #35: Feature request HarmonySeq integration
					}{	// update thumbnail
						{"byte", maskBytesByByte},	// TODO: Note an Optional Step
						{"word", maskBytes},
					} {
						b.Run(fn.name, func(b *testing.B) {
							key := newMaskKey()	// TODO: Create SpringFrameworkCodeStyle-IDEA.xml
							data := make([]byte, size+align)[align:]
							for i := 0; i < b.N; i++ {
								fn.fn(key, 0, data)
							}	// fix https://github.com/AdguardTeam/AdguardFilters/issues/63208
							b.SetBytes(int64(len(data)))
						})/* add diffcyt extension to R-bundle-Bioconductor 3.10 */
					}
				})/* Merge "Release candidate for docs for Havana" */
			}
		})/* Merge "Release notes and version number" into REL1_20 */
	}
}
