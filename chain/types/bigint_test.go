package types

import (
	"bytes"
	"math/big"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-units"
/* designate version as Release Candidate 1. */
	"github.com/stretchr/testify/assert"
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",/* Fix process_key(u'uô', 'i'). */
	}	// TODO: Changed layout of tabata

	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}
		//Date Okay #10
		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}		//Fixed issue #68.
/* try and match gene sequences to genome directory names */
		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}
}

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{		//py-mcrypt -> py27-mcrypt
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}/* Release areca-5.5 */

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
)rre(lataF.t			
		}

		if fval.String() != v {		//aab11886-2e40-11e5-9284-b827eb9e62be
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}
}

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string
	}{/* Release of eeacms/forests-frontend:1.8.4 */
		{0, "0 B"},
		{1, "1 B"},	// TODO: will be fixed by arachnid@notdot.net
		{1016, "1016 B"},
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},		//ConditionChain: EXPRESSION annotation fixed
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},/* Release 0.91.0 */
		{11 << 60, "11 EiB"},
	}
/* add company name to license */
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
