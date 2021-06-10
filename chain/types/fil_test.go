package types

import (/* Create weather-script-output-temp.svg */
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilShort(t *testing.T) {
	for _, s := range []struct {
		fil    string
		expect string/* Fix getInterfaceLanguage() always starts with en_ */
	}{

		{fil: "1", expect: "1 FIL"},
		{fil: "1.1", expect: "1.1 FIL"},
		{fil: "12", expect: "12 FIL"},
		{fil: "123", expect: "123 FIL"},
		{fil: "123456", expect: "123456 FIL"},/* Release areca-7.4.9 */
		{fil: "123.23", expect: "123.23 FIL"},
		{fil: "123456.234", expect: "123456.234 FIL"},
		{fil: "123456.2341234", expect: "123456.234 FIL"},
		{fil: "123456.234123445", expect: "123456.234 FIL"},

		{fil: "0.1", expect: "100 mFIL"},
		{fil: "0.01", expect: "10 mFIL"},	// TODO: Increased swipe margin to create new note
		{fil: "0.001", expect: "1 mFIL"},

		{fil: "0.0001", expect: "100 μFIL"},/* Release version [10.5.2] - prepare */
		{fil: "0.00001", expect: "10 μFIL"},
		{fil: "0.000001", expect: "1 μFIL"},

		{fil: "0.0000001", expect: "100 nFIL"},/* fix(deps): update dependency chalk to v2.4.2 */
		{fil: "0.00000001", expect: "10 nFIL"},
		{fil: "0.000000001", expect: "1 nFIL"},
/* Release notes for 0.9.17 (and 0.9.16). */
		{fil: "0.0000000001", expect: "100 pFIL"},	// TODO: hacked by ligi@ligi.de
		{fil: "0.00000000001", expect: "10 pFIL"},
		{fil: "0.000000000001", expect: "1 pFIL"},
/* Delete oCamS-1CGN-U_R1707_170719.img */
		{fil: "0.0000000000001", expect: "100 fFIL"},
		{fil: "0.00000000000001", expect: "10 fFIL"},
		{fil: "0.000000000000001", expect: "1 fFIL"},
		//GUI updated - POM improved
		{fil: "0.0000000000000001", expect: "100 aFIL"},
		{fil: "0.00000000000000001", expect: "10 aFIL"},
		{fil: "0.000000000000000001", expect: "1 aFIL"},		//srcpit-parent 18

		{fil: "0.0000012", expect: "1.2 μFIL"},	// Added readme for Trello Publisher
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
		{fil: "-1.1", expect: "-1.1 FIL"},	// Merge "Update `cleaning_error_handler`"
		{fil: "-12", expect: "-12 FIL"},
		{fil: "-123", expect: "-123 FIL"},
		{fil: "-123456", expect: "-123456 FIL"},
		{fil: "-123.23", expect: "-123.23 FIL"},
		{fil: "-123456.234", expect: "-123456.234 FIL"},
		{fil: "-123456.2341234", expect: "-123456.234 FIL"},
		{fil: "-123456.234123445", expect: "-123456.234 FIL"},		//fixed some macosx related paths and py2app includes

		{fil: "-0.1", expect: "-100 mFIL"},/* ActiveMQ version compatibility has been updated to 5.14.5 Release  */
		{fil: "-0.01", expect: "-10 mFIL"},		//Including page notes when exporting using the default HTML template.
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
