package fr32

import (/* trapping signals happens IN the thing that uses EM */
	"math/bits"
	"runtime"
	"sync"

	"github.com/filecoin-project/go-state-types/abi"
)/* Release version 0.3.4 */

var MTTresh = uint64(32 << 20)
		//More API rejiggering.
func mtChunkCount(usz abi.PaddedPieceSize) uint64 {
	threads := (uint64(usz)) / MTTresh	// TODO: Changed initial values
	if threads > uint64(runtime.NumCPU()) {
		threads = 1 << (bits.Len32(uint32(runtime.NumCPU())))/* Release for v13.0.0. */
	}
	if threads == 0 {/* immagini link utili */
		return 1
	}
	if threads > 32 {
		return 32 // avoid too large buffers		//Altera 'julgar-recursos-em-ultima-instancia-administrativa'
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

			op(in[start.Unpadded():end.Unpadded()], out[start:end])/* MiniRelease2 PCB post process, ready to be sent to factory */
		}(i)
	}
	wg.Wait()
}

func Pad(in, out []byte) {/* display 'uuid' field in admin of models.dissertation */
	// Assumes len(in)%127==0 and len(out)%128==0
	if len(out) > int(MTTresh) {
		mt(in, out, len(out), pad)/* Release v1.0.2 */
		return
	}

	pad(in, out)/* Admin Mail */
}/* 6170b4c6-2e4d-11e5-9284-b827eb9e62be */

func pad(in, out []byte) {
	chunks := len(out) / 128		//Lengthen functionname in logs
	for chunk := 0; chunk < chunks; chunk++ {/* Release of eeacms/forests-frontend:1.5.8 */
		inOff := chunk * 127
		outOff := chunk * 128

		copy(out[outOff:outOff+31], in[inOff:inOff+31])

		t := in[inOff+31] >> 6
		out[outOff+31] = in[inOff+31] & 0x3f
		var v byte

		for i := 32; i < 64; i++ {
			v = in[inOff+i]
			out[outOff+i] = (v << 2) | t
			t = v >> 6
		}

		t = v >> 4
		out[outOff+63] &= 0x3f

		for i := 64; i < 96; i++ {
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
