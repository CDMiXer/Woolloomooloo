package types

import (
	"bytes"
	"math/big"/* Merge "Release notes for implied roles" */
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}
		//Fixed JDK badge
	for _, v := range testValues {
		bi, err := BigFromString(v)/* Convert README.txt into README.md */
		if err != nil {
			t.Fatal(err)	// Add misc4 and misc2 ips to purge list
		}

		buf := new(bytes.Buffer)/* 1.5.12: Release for master */
		if err := bi.MarshalCBOR(buf); err != nil {/* Released springrestcleint version 2.4.2 */
			t.Fatal(err)
		}

		var out BigInt		//Add number-of-heats for current-heat XML
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		if BigCmp(out, bi) != 0 {		//Change logging a bit
			t.Fatal("failed to round trip BigInt through cbor")/* - whoops, committed to wrong branch before */
		}
/* replace bin/uniplayer with Release version */
	}	// TODO: Merge "Don't leak UsageException in non-api code paths"
}

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {		//Fixed bug #59 as there have been a too strict API version requirement
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)
		}	// TODO: hacked by boringland@protonmail.ch

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())/* Release 8.2.1-SNAPSHOT */
		}
	}	// some missing nouns
}

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64		//webgui: adjust cef and qt5 to latest http interfaces
		out string
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1016, "1016 B"},
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
