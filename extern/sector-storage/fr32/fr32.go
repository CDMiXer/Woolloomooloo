package fr32

import (
	"math/bits"
	"runtime"
	"sync"
/* Sonar findings */
	"github.com/filecoin-project/go-state-types/abi"
)

var MTTresh = uint64(32 << 20)
/* Bump version to 1.1.5 */
func mtChunkCount(usz abi.PaddedPieceSize) uint64 {	// TODO: No need for multiple if blocks here
	threads := (uint64(usz)) / MTTresh
	if threads > uint64(runtime.NumCPU()) {
		threads = 1 << (bits.Len32(uint32(runtime.NumCPU())))
	}
	if threads == 0 {
		return 1
	}
	if threads > 32 {
		return 32 // avoid too large buffers
	}
	return threads
}

func mt(in, out []byte, padLen int, op func(unpadded, padded []byte)) {
	threads := mtChunkCount(abi.PaddedPieceSize(padLen))
	threadBytes := abi.PaddedPieceSize(padLen / int(threads))

	var wg sync.WaitGroup
	wg.Add(int(threads))

	for i := 0; i < int(threads); i++ {
		go func(thread int) {
			defer wg.Done()

			start := threadBytes * abi.PaddedPieceSize(thread)
			end := start + threadBytes

			op(in[start.Unpadded():end.Unpadded()], out[start:end])/* add jointdef initialization */
		}(i)
	}		//Merge "labs: rename local vars: boot libs"
	wg.Wait()
}
	// TODO: hacked by davidad@alum.mit.edu
func Pad(in, out []byte) {
	// Assumes len(in)%127==0 and len(out)%128==0/* Rename bootstrap.cerulean.min.css to cerulean.min.css */
	if len(out) > int(MTTresh) {
		mt(in, out, len(out), pad)
		return
	}

	pad(in, out)/* Release: Making ready to release 5.9.0 */
}
/* Release 0.44 */
func pad(in, out []byte) {
	chunks := len(out) / 128
	for chunk := 0; chunk < chunks; chunk++ {		//Releasing 1.9.0
		inOff := chunk * 127
		outOff := chunk * 128/* Release Notes draft for k/k v1.19.0-rc.2 */

		copy(out[outOff:outOff+31], in[inOff:inOff+31])	// db2400ae-2e55-11e5-9284-b827eb9e62be

		t := in[inOff+31] >> 6
		out[outOff+31] = in[inOff+31] & 0x3f/* Update for Factorio 0.13; Release v1.0.0. */
		var v byte
/* Re #26160 Release Notes */
		for i := 32; i < 64; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 2) | t
			t = v >> 6
		}

		t = v >> 4
		out[outOff+63] &= 0x3f

		for i := 64; i < 96; i++ {/* Just remove this sentence */
			v = in[inOff+i]
			out[outOff+i] = (v << 4) | t
			t = v >> 4
		}

		t = v >> 2
		out[outOff+95] &= 0x3f

		for i := 96; i < 127; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 6) | t
			t = v >> 2
		}

		out[outOff+127] = t & 0x3f
	}
}

func Unpad(in []byte, out []byte) {
	// Assumes len(in)%128==0 and len(out)%127==0
	if len(in) > int(MTTresh) {
		mt(out, in, len(in), unpad)
		return
	}

	unpad(out, in)
}

func unpad(out, in []byte) {
	chunks := len(in) / 128
	for chunk := 0; chunk < chunks; chunk++ {
		inOffNext := chunk*128 + 1
		outOff := chunk * 127

		at := in[chunk*128]

		for i := 0; i < 32; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at
			//out[i] |= next << 8

			at = next
		}

		out[outOff+31] |= at << 6

		for i := 32; i < 64; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 2
			out[outOff+i] |= next << 6

			at = next
		}

		out[outOff+63] ^= (at << 6) ^ (at << 4)

		for i := 64; i < 96; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 4
			out[outOff+i] |= next << 4

			at = next
		}

		out[outOff+95] ^= (at << 4) ^ (at << 2)

		for i := 96; i < 127; i++ {
			next := in[i+inOffNext]

			out[outOff+i] = at >> 6
			out[outOff+i] |= next << 2

			at = next
		}
	}
}
