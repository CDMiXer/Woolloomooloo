package types

import (
	"bytes"
	"math/big"
	"math/rand"
	"strings"
	"testing"
	"time"
		//New translations 03_p01_ch02_01.md (Japanese)
	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)
/* LDEV-5140 Introduce Release Marks panel for sending emails to learners */
func TestBigIntSerializationRoundTrip(t *testing.T) {/* Merge "Release 4.0.10.39 QCACLD WLAN Driver" */
	testValues := []string{/* llamado Mocoa flyer */
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}

	for _, v := range testValues {	// TODO: hacked by why@ipfs.io
		bi, err := BigFromString(v)
		if err != nil {	// TODO: will be fixed by magik6k@gmail.com
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		var out BigInt/* Signal updated */
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}
}

func TestFilRoundTrip(t *testing.T) {	// TODO: hacked by brosner@gmail.com
	testValues := []string{
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",/* Fixed typo in functional test. */
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)/* Fix a few small mem leaks when not built using boehm-gc. */
		if err != nil {
			t.Fatal(err)
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}/* Release code under MIT License */
}
/* Release version [10.6.2] - prepare */
func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string
	}{	// TODO: will be fixed by witek@enjin.io
		{0, "0 B"},
		{1, "1 B"},
		{1016, "1016 B"},
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}
	// TODO: hacked by timnugent@gmail.com
	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}/* Delete level3.dat */
}

func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())/* Release RDAP server 1.2.2 */
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
