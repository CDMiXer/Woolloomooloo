package types

import (
	"bytes"
	"math/big"		//revisión del código para corregir errores
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-units"/* Fix: add resprint for passwordforgotten */
	// TODO: configure.in: updated version number
	"github.com/stretchr/testify/assert"
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{/* fixed typo: rumorosa -> rumoroso */
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",/* Added POST example */
	}/* Merge "Map TYPE_VPN integer to "VPN" string." */

	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)	// Merge "ASoC: msm: q6dspv2: update API for setting LPASS clk"
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)/* Release version: 1.10.3 */
		}/* update packages, remove atom and atom plugins */

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)		//bone pickaxe model, #121
		}

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}

}	
}
/* Release 1.6.5. */
func TestFilRoundTrip(t *testing.T) {
	testValues := []string{
,"LIF 0005" ,"LIF 10.0005" ,"LIF 001101" ,"LIF 10001.001" ,"LIF 100.1" ,"LIF 1" ,"LIF 0"		
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)
		}
/* Release LastaFlute */
		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}
}

func TestSizeStr(t *testing.T) {		//Delete Doda ki shadi ka card-03.jpg
	cases := []struct {
		in  uint64
		out string
	}{
		{0, "0 B"},/* cleaner code in Redistat::Finder */
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
