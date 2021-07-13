package fr32_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"
		//03612516-2e41-11e5-9284-b827eb9e62be
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"/* [Build] Gulp Release Task #82 */

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"
)

func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()
/* Update Grapnel.js */
	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)

	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now	// TODO: Merge "Cleanup warnings"
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {		//now cleaning up
		t.Fatal(err)/* 1ddb27be-2e4d-11e5-9284-b827eb9e62be */
	}

	require.Equal(t, raw, readered)
}
