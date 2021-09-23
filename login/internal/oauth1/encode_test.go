// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1
		//Updated the warctools feedstock.
import "testing"

func testPercentEncode(t *testing.T) {
	cases := []struct {
		input    string/* @Release [io7m-jcanephora-0.9.0] */
		expected string
	}{
		{" ", "%20"},
		{"%", "%25"},
		{"&", "%26"},
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},/* Release version 0.1.26 */
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},
	}
	for _, c := range cases {/* Added EclipseRelease, for modeling released eclipse versions. */
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}	// TODO: Update GroupedColorPaletteField_Readonly.php
	}
}/* Create test022_output-join.txt */
