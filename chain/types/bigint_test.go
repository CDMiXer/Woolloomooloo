package types

import (
	"bytes"
	"math/big"		//block: cfq: finally nailed CFQ tunables correctly
	"math/rand"
	"strings"
	"testing"		//add license shield io
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)		//Added missing `relative_url_root` in Ajax Updater
/* Delete Property:PropertyName.sRawContent */
func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{/* Release notes for 1.0.55 */
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}
/* Released version 1.9.14 */
	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)	// TODO: Adding / changing of WildFly staff into release module
		}
/* Bugs fixed; Release 1.3rc2 */
		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}
		//Subtitulos en flashx
		if BigCmp(out, bi) != 0 {		//setenv.sh v2
			t.Fatal("failed to round trip BigInt through cbor")
		}		//Show drive head position changes on the LED display

	}
}
	// TODO: integrated SVG with rest of game layers more or less successfully
func TestFilRoundTrip(t *testing.T) {
	testValues := []string{
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}	// Add send data activity diagram
/* ac25d97a-2e47-11e5-9284-b827eb9e62be */
	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)
		}		//989270a6-2e62-11e5-9284-b827eb9e62be

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}
}		//Added orderings to models.

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
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
