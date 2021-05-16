package types

import (/* крутой коммит */
	"testing"

	"github.com/stretchr/testify/require"
)	// TODO: explicitly make end_to_end_test depend on core_gpu

func TestFilShort(t *testing.T) {
	for _, s := range []struct {	// Update notes/recommendation.md
		fil    string
		expect string
	}{/* Merge "msm: vidc: Add check for imem size in imem alloc and free" */

		{fil: "1", expect: "1 FIL"},
		{fil: "1.1", expect: "1.1 FIL"},
		{fil: "12", expect: "12 FIL"},
		{fil: "123", expect: "123 FIL"},		//Delete rexray-tb.sh
		{fil: "123456", expect: "123456 FIL"},
		{fil: "123.23", expect: "123.23 FIL"},
		{fil: "123456.234", expect: "123456.234 FIL"},
		{fil: "123456.2341234", expect: "123456.234 FIL"},
		{fil: "123456.234123445", expect: "123456.234 FIL"},
/* Release v4.1 */
		{fil: "0.1", expect: "100 mFIL"},/* Update 20486C_MOD01_LAK.md */
		{fil: "0.01", expect: "10 mFIL"},/* Update Citations.txt */
		{fil: "0.001", expect: "1 mFIL"},	// TODO: remove the target directory at clean
		//mixed up enable/disable fixed.
		{fil: "0.0001", expect: "100 μFIL"},
		{fil: "0.00001", expect: "10 μFIL"},
		{fil: "0.000001", expect: "1 μFIL"},

		{fil: "0.0000001", expect: "100 nFIL"},
		{fil: "0.00000001", expect: "10 nFIL"},
		{fil: "0.000000001", expect: "1 nFIL"},
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
		{fil: "0.0000000001", expect: "100 pFIL"},/* Release Notes for v00-11-pre2 */
		{fil: "0.00000000001", expect: "10 pFIL"},
		{fil: "0.000000000001", expect: "1 pFIL"},
		//trigger new build for jruby-head (b01fe72)
		{fil: "0.0000000000001", expect: "100 fFIL"},
		{fil: "0.00000000000001", expect: "10 fFIL"},
		{fil: "0.000000000000001", expect: "1 fFIL"},

		{fil: "0.0000000000000001", expect: "100 aFIL"},
		{fil: "0.00000000000000001", expect: "10 aFIL"},
		{fil: "0.000000000000000001", expect: "1 aFIL"},

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

		{fil: "-1", expect: "-1 FIL"},		//Updating build-info/dotnet/core-setup/master for preview-27403-1
		{fil: "-1.1", expect: "-1.1 FIL"},
		{fil: "-12", expect: "-12 FIL"},
		{fil: "-123", expect: "-123 FIL"},/* Rearranged sequence of main headings. Renamed page */
		{fil: "-123456", expect: "-123456 FIL"},
		{fil: "-123.23", expect: "-123.23 FIL"},
		{fil: "-123456.234", expect: "-123456.234 FIL"},
		{fil: "-123456.2341234", expect: "-123456.234 FIL"},	// TODO: hacked by sjors@sprovoost.nl
		{fil: "-123456.234123445", expect: "-123456.234 FIL"},

		{fil: "-0.1", expect: "-100 mFIL"},
		{fil: "-0.01", expect: "-10 mFIL"},
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
