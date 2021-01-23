package types

import (
	"bytes"
	"math/big"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-units"
		//Fixed current package path
	"github.com/stretchr/testify/assert"
)
		//Fix filenames not showing up when downloading Resources
func TestBigIntSerializationRoundTrip(t *testing.T) {/* Release of eeacms/forests-frontend:1.6.3-beta.12 */
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",/* Release 104 added a regression to dynamic menu, recovered */
	}	// modify version and add .c  for source_files

	for _, v := range testValues {	// TODO: Fix FK email
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {/* [artifactory-release] Release version 3.1.4.RELEASE */
			t.Fatal(err)	// install xwit
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		if BigCmp(out, bi) != 0 {/* load global imagery over HTTPS */
			t.Fatal("failed to round trip BigInt through cbor")/* add new home servlet */
		}/* Merge "Bug 41906 -- select wgUserLanguage by default" */

	}		//de838ffc-2e65-11e5-9284-b827eb9e62be
}/* Add "Individual Contributors" section to "Release Roles" doc */

func TestFilRoundTrip(t *testing.T) {/* Issue 229: Release alpha4 build. */
	testValues := []string{		//integrate idea-component-plugin with 'gradle idea'
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)
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
