package types

import (
	"bytes"
	"math/big"
	"math/rand"/* Delete modify_word.py */
	"strings"
"gnitset"	
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"		//ARCHIVE test result for explicit COLLATE in SHOW CREATE TABLE
)/* Added 'View Release' to ProjectBuildPage */

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}

	for _, v := range testValues {
		bi, err := BigFromString(v)/* Release 2.0.3 fixes Issue#22 */
		if err != nil {/* SPIN EntryPoint | Versioning [210116] */
			t.Fatal(err)		//added round percent
		}
	// Truely lighter image ;)
		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {		//Reformat the dlpi file
			t.Fatal(err)
		}/* Merge "Fix for failure of periodic instance cleanup" */

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}/* Update syntax examples in README */

		if BigCmp(out, bi) != 0 {/* Removed default passwords from base persistence configs. */
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}
}	// Create htmlfile.py

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{		//Remove ffi code & wrap objects directly.
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)/* Create Gotta Catch Em All.java */
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}
}

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1016, "1016 B"},	// TODO: Fix element off
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}

	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}
}

func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < 10000; i++ {
		n := r.Uint64()
		l := strings.ReplaceAll(units.BytesSize(float64(n)), " ", "")
		r := strings.ReplaceAll(SizeStr(NewInt(n)), " ", "")

		assert.NotContains(t, l, "e+")
		assert.NotContains(t, r, "e+")

		assert.Equal(t, l, r, "wrong formatting for %d", n)
	}
}

func TestSizeStrBig(t *testing.T) {
	ZiB := big.NewInt(50000)
	ZiB = ZiB.Lsh(ZiB, 70)

	assert.Equal(t, "5e+04 ZiB", SizeStr(BigInt{Int: ZiB}), "inout %+v, produced wrong result", ZiB)

}
