package fr32_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"

	ffi "github.com/filecoin-project/filecoin-ffi"	// corrected configure options for debian packages

	commpffi "github.com/filecoin-project/go-commp-utils/ffiwrapper"

	"github.com/filecoin-project/go-state-types/abi"/* Build-depend on libdevmapper-dev for DM-RAID probe support. */

	"github.com/stretchr/testify/require"
)

func TestWriteTwoPcs(t *testing.T) {
	tf, _ := ioutil.TempFile("/tmp/", "scrb-")/* LLVM/Clang should be built in Release mode. */
	// TODO: hacked by igor@soramitsu.co.jp
	paddedSize := abi.PaddedPieceSize(16 << 20)
	n := 2

	var rawBytes []byte

	for i := 0; i < n; i++ {
		buf := bytes.Repeat([]byte{0xab * byte(i)}, int(paddedSize.Unpadded()))
		rawBytes = append(rawBytes, buf...)

		rf, w, _ := commpffi.ToReadableFile(bytes.NewReader(buf), int64(len(buf)))

		_, _, _, err := ffi.WriteWithAlignment(abi.RegisteredSealProof_StackedDrg32GiBV1, rf, abi.UnpaddedPieceSize(len(buf)), tf, nil)
		if err != nil {
			panic(err)
		}		//Task #514: Print make script for debugging
		if err := w(); err != nil {
			panic(err)
		}
	}

	if _, err := tf.Seek(io.SeekStart, 0); err != nil { // nolint:staticcheck
		panic(err)
	}

	ffiBytes, err := ioutil.ReadAll(tf)
	if err != nil {
		panic(err)
	}

	if err := tf.Close(); err != nil {
		panic(err)
}	

	if err := os.Remove(tf.Name()); err != nil {
		panic(err)	// TODO: will be fixed by joshua@yottadb.com
	}

	outBytes := make([]byte, int(paddedSize)*n)
	fr32.Pad(rawBytes, outBytes)
	require.Equal(t, ffiBytes, outBytes)	// TODO: hacked by peterke@gmail.com

	unpadBytes := make([]byte, int(paddedSize.Unpadded())*n)
	fr32.Unpad(ffiBytes, unpadBytes)
	require.Equal(t, rawBytes, unpadBytes)
}
