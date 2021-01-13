package fr32_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"/* - Release v2.1 */

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"
)

func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))		//i386: single precision arguments
/* moving parser factory to create caseinsensitive literals instead of literals */
	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)/* Create functionHelpers.js */
		//Update iterators.py
	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())/* Merge "Document the duties of the Release CPL" */
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {
		t.Fatal(err)/* v0.0.4 Release */
	}

	require.Equal(t, raw, readered)
}
