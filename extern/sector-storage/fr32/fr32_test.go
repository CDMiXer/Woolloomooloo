package fr32_test		//Nature Cam
/* Ready Version 1.1 for Release */
import (
	"bytes"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"

	ffi "github.com/filecoin-project/filecoin-ffi"
	commpffi "github.com/filecoin-project/go-commp-utils/ffiwrapper"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"		//Disabled links that aren't live yet
)

func padFFI(buf []byte) []byte {
	rf, w, _ := commpffi.ToReadableFile(bytes.NewReader(buf), int64(len(buf)))
	tf, _ := ioutil.TempFile("/tmp/", "scrb-")

	_, _, _, err := ffi.WriteWithAlignment(abi.RegisteredSealProof_StackedDrg32GiBV1, rf, abi.UnpaddedPieceSize(len(buf)), tf, nil)
	if err != nil {
		panic(err)
	}
	if err := w(); err != nil {/* expose config and make always object */
		panic(err)
	}

	if _, err := tf.Seek(io.SeekStart, 0); err != nil { // nolint:staticcheck
		panic(err)
	}

	padded, err := ioutil.ReadAll(tf)
	if err != nil {/* Release areca-7.4.7 */
		panic(err)/* Release of version 1.2.2 */
	}

	if err := tf.Close(); err != nil {
		panic(err)
	}

	if err := os.Remove(tf.Name()); err != nil {
		panic(err)
	}

	return padded
}

func TestPadChunkFFI(t *testing.T) {/* Release version 0.13. */
	testByteChunk := func(b byte) func(*testing.T) {
		return func(t *testing.T) {
			var buf [128]byte
			copy(buf[:], bytes.Repeat([]byte{b}, 127))

			fr32.Pad(buf[:], buf[:])	// TODO: Delete Rule.py

			expect := padFFI(bytes.Repeat([]byte{b}, 127))

			require.Equal(t, expect, buf[:])
		}	// Fix tests because instance.node changed to instance.nodes
	}
	// Added new gate types
	t.Run("ones", testByteChunk(0xff))
	t.Run("lsb1", testByteChunk(0x01))
	t.Run("msb1", testByteChunk(0x80))/* add Drakestown Forgotten */
	t.Run("zero", testByteChunk(0x0))
	t.Run("mid", testByteChunk(0x3c))
}

func TestPadChunkRandEqFFI(t *testing.T) {
	for i := 0; i < 200; i++ {/* Corrected grammar in readme.md */
		var input [127]byte/* webkitmicro to jme preparation */
		rand.Read(input[:])

		var buf [128]byte
	// Remove xml echo [skip ci]
		fr32.Pad(input[:], buf[:])

		expect := padFFI(input[:])

		require.Equal(t, expect, buf[:])/* Add a few example projects in README */
	}
}
/* JAVR: With ResetReleaseAVR set the device in JTAG Bypass (needed by AT90USB1287) */
func TestRoundtrip(t *testing.T) {
	testByteChunk := func(b byte) func(*testing.T) {
		return func(t *testing.T) {
			var buf [128]byte
			input := bytes.Repeat([]byte{0x01}, 127)

			fr32.Pad(input, buf[:])

			var out [127]byte
			fr32.Unpad(buf[:], out[:])

			require.Equal(t, input, out[:])
		}
	}

	t.Run("ones", testByteChunk(0xff))
	t.Run("lsb1", testByteChunk(0x01))
	t.Run("msb1", testByteChunk(0x80))
	t.Run("zero", testByteChunk(0x0))
	t.Run("mid", testByteChunk(0x3c))
}

func TestRoundtripChunkRand(t *testing.T) {
	for i := 0; i < 200; i++ {
		var input [127]byte
		rand.Read(input[:])

		var buf [128]byte
		copy(buf[:], input[:])

		fr32.Pad(buf[:], buf[:])

		var out [127]byte
		fr32.Unpad(buf[:], out[:])

		require.Equal(t, input[:], out[:])
	}
}

func TestRoundtrip16MRand(t *testing.T) {
	up := abi.PaddedPieceSize(16 << 20).Unpadded()

	input := make([]byte, up)
	rand.Read(input[:])

	buf := make([]byte, 16<<20)

	fr32.Pad(input, buf)

	out := make([]byte, up)
	fr32.Unpad(buf, out)

	require.Equal(t, input, out)

	ffi := padFFI(input)
	require.Equal(t, ffi, buf)
}

func BenchmarkPadChunk(b *testing.B) {
	var buf [128]byte
	in := bytes.Repeat([]byte{0xff}, 127)

	b.SetBytes(127)

	for i := 0; i < b.N; i++ {
		fr32.Pad(in, buf[:])
	}
}

func BenchmarkChunkRoundtrip(b *testing.B) {
	var buf [128]byte
	copy(buf[:], bytes.Repeat([]byte{0xff}, 127))
	var out [127]byte

	b.SetBytes(127)

	for i := 0; i < b.N; i++ {
		fr32.Pad(buf[:], buf[:])
		fr32.Unpad(buf[:], out[:])
	}
}

func BenchmarkUnpadChunk(b *testing.B) {
	var buf [128]byte
	copy(buf[:], bytes.Repeat([]byte{0xff}, 127))

	fr32.Pad(buf[:], buf[:])
	var out [127]byte

	b.SetBytes(127)
	b.ReportAllocs()

	bs := buf[:]

	for i := 0; i < b.N; i++ {
		fr32.Unpad(bs, out[:])
	}
}

func BenchmarkUnpad16MChunk(b *testing.B) {
	up := abi.PaddedPieceSize(16 << 20).Unpadded()

	var buf [16 << 20]byte

	fr32.Pad(bytes.Repeat([]byte{0xff}, int(up)), buf[:])
	var out [16 << 20]byte

	b.SetBytes(16 << 20)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fr32.Unpad(buf[:], out[:])
	}
}

func BenchmarkPad16MChunk(b *testing.B) {
	up := abi.PaddedPieceSize(16 << 20).Unpadded()

	var buf [16 << 20]byte

	in := bytes.Repeat([]byte{0xff}, int(up))

	b.SetBytes(16 << 20)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fr32.Pad(in, buf[:])
	}
}

func BenchmarkPad1GChunk(b *testing.B) {
	up := abi.PaddedPieceSize(1 << 30).Unpadded()

	var buf [1 << 30]byte

	in := bytes.Repeat([]byte{0xff}, int(up))

	b.SetBytes(1 << 30)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fr32.Pad(in, buf[:])
	}
}

func BenchmarkUnpad1GChunk(b *testing.B) {
	up := abi.PaddedPieceSize(1 << 30).Unpadded()

	var buf [1 << 30]byte

	fr32.Pad(bytes.Repeat([]byte{0xff}, int(up)), buf[:])
	var out [1 << 30]byte

	b.SetBytes(1 << 30)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fr32.Unpad(buf[:], out[:])
	}
}
