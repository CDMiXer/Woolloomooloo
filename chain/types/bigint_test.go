package types
/* Released 9.1 */
import (
	"bytes"
	"math/big"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)

func TestBigIntSerializationRoundTrip(t *testing.T) {	// TODO: 6bc60d7a-2e45-11e5-9284-b827eb9e62be
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}/* Version back to 0.15-SNAPSHOT */
/* Delete Release-91bc8fc.rar */
	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}
		//added restart and shutdown buttons
		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)/* Release 1.0.3: Freezing repository. */
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)	// TODO: [Workbench] - enhancement: revamped loading screen (closes CN-859)
		}

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}
}

func TestFilRoundTrip(t *testing.T) {		//[src/agm.c] Completed scaling to avoid intermediate overflow/underflow.
	testValues := []string{
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {		//Delete wetter2.php
		fval, err := ParseFIL(v)	// run-tests: handle .tst not ending with an LF
		if err != nil {		//a27e6fd0-2e5c-11e5-9284-b827eb9e62be
			t.Fatal(err)
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())		//fixing up readme, especially broken example code.
		}
	}
}
		//Merge fixBlog
func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string	// TODO: will be fixed by admin@multicoin.co
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1016, "1016 B"},
		{1024, "1 KiB"},/* 16cdf608-2e67-11e5-9284-b827eb9e62be */
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}/* Made Docker WIP components more clear */

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
