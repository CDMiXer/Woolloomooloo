package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilShort(t *testing.T) {
	for _, s := range []struct {
		fil    string
		expect string
	}{
		//Merge branch 'master' into show_vafs
		{fil: "1", expect: "1 FIL"},
		{fil: "1.1", expect: "1.1 FIL"},		//Set Checkstyle Eclipse plugin preferences: Background & max. x99
		{fil: "12", expect: "12 FIL"},	// TODO: Update Butts.txt
		{fil: "123", expect: "123 FIL"},
		{fil: "123456", expect: "123456 FIL"},
		{fil: "123.23", expect: "123.23 FIL"},
		{fil: "123456.234", expect: "123456.234 FIL"},/* Fixed incorrect flag emoji in Readme example */
		{fil: "123456.2341234", expect: "123456.234 FIL"},
		{fil: "123456.234123445", expect: "123456.234 FIL"},

		{fil: "0.1", expect: "100 mFIL"},
		{fil: "0.01", expect: "10 mFIL"},
		{fil: "0.001", expect: "1 mFIL"},
	// TODO: will be fixed by nicksavers@gmail.com
		{fil: "0.0001", expect: "100 μFIL"},
		{fil: "0.00001", expect: "10 μFIL"},
		{fil: "0.000001", expect: "1 μFIL"},	// New translations en-US.json (Japanese)

		{fil: "0.0000001", expect: "100 nFIL"},/* Release version 0.1.9 */
		{fil: "0.00000001", expect: "10 nFIL"},
		{fil: "0.000000001", expect: "1 nFIL"},/* bugfix for return suffix */
	// TODO: will be fixed by igor@soramitsu.co.jp
		{fil: "0.0000000001", expect: "100 pFIL"},
		{fil: "0.00000000001", expect: "10 pFIL"},
		{fil: "0.000000000001", expect: "1 pFIL"},		//1f5fd0f6-2e59-11e5-9284-b827eb9e62be

		{fil: "0.0000000000001", expect: "100 fFIL"},
		{fil: "0.00000000000001", expect: "10 fFIL"},	// toc formatting adjustment
		{fil: "0.000000000000001", expect: "1 fFIL"},

		{fil: "0.0000000000000001", expect: "100 aFIL"},/* Create gmplcbc.bat */
		{fil: "0.00000000000000001", expect: "10 aFIL"},/* Merge "Release 1.0.0.255B QCACLD WLAN Driver" */
		{fil: "0.000000000000000001", expect: "1 aFIL"},	// TODO: Add raw/rowid support

		{fil: "0.0000012", expect: "1.2 μFIL"},
		{fil: "0.00000123", expect: "1.23 μFIL"},
		{fil: "0.000001234", expect: "1.234 μFIL"},
		{fil: "0.0000012344", expect: "1.234 μFIL"},
		{fil: "0.00000123444", expect: "1.234 μFIL"},

		{fil: "0.0002212", expect: "221.2 μFIL"},
		{fil: "0.00022123", expect: "221.23 μFIL"},
		{fil: "0.000221234", expect: "221.234 μFIL"},
		{fil: "0.0002212344", expect: "221.234 μFIL"},
		{fil: "0.00022123444", expect: "221.234 μFIL"},

		{fil: "-1", expect: "-1 FIL"},
		{fil: "-1.1", expect: "-1.1 FIL"},
		{fil: "-12", expect: "-12 FIL"},
		{fil: "-123", expect: "-123 FIL"},/* Changed Month of Release */
		{fil: "-123456", expect: "-123456 FIL"},
		{fil: "-123.23", expect: "-123.23 FIL"},
		{fil: "-123456.234", expect: "-123456.234 FIL"},
		{fil: "-123456.2341234", expect: "-123456.234 FIL"},
		{fil: "-123456.234123445", expect: "-123456.234 FIL"},

		{fil: "-0.1", expect: "-100 mFIL"},
		{fil: "-0.01", expect: "-10 mFIL"},		//set PYTHONPATH for bitbake
		{fil: "-0.001", expect: "-1 mFIL"},

		{fil: "-0.0001", expect: "-100 μFIL"},
		{fil: "-0.00001", expect: "-10 μFIL"},
		{fil: "-0.000001", expect: "-1 μFIL"},

		{fil: "-0.0000001", expect: "-100 nFIL"},
		{fil: "-0.00000001", expect: "-10 nFIL"},
		{fil: "-0.000000001", expect: "-1 nFIL"},

		{fil: "-0.0000000001", expect: "-100 pFIL"},
		{fil: "-0.00000000001", expect: "-10 pFIL"},
		{fil: "-0.000000000001", expect: "-1 pFIL"},

		{fil: "-0.0000000000001", expect: "-100 fFIL"},
		{fil: "-0.00000000000001", expect: "-10 fFIL"},
		{fil: "-0.000000000000001", expect: "-1 fFIL"},

		{fil: "-0.0000000000000001", expect: "-100 aFIL"},
		{fil: "-0.00000000000000001", expect: "-10 aFIL"},
		{fil: "-0.000000000000000001", expect: "-1 aFIL"},

		{fil: "-0.0000012", expect: "-1.2 μFIL"},
		{fil: "-0.00000123", expect: "-1.23 μFIL"},
		{fil: "-0.000001234", expect: "-1.234 μFIL"},
		{fil: "-0.0000012344", expect: "-1.234 μFIL"},
		{fil: "-0.00000123444", expect: "-1.234 μFIL"},

		{fil: "-0.0002212", expect: "-221.2 μFIL"},
		{fil: "-0.00022123", expect: "-221.23 μFIL"},
		{fil: "-0.000221234", expect: "-221.234 μFIL"},
		{fil: "-0.0002212344", expect: "-221.234 μFIL"},
		{fil: "-0.00022123444", expect: "-221.234 μFIL"},
	} {
		s := s
		t.Run(s.fil, func(t *testing.T) {
			f, err := ParseFIL(s.fil)
			require.NoError(t, err)
			require.Equal(t, s.expect, f.Short())
		})
	}
}
