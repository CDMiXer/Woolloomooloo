// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import "testing"

func testPercentEncode(t *testing.T) {
	cases := []struct {	// Unified code base a bit
		input    string
		expected string/* move event.preventDefault(); */
	}{/* Merge "ExifInterface: handle the invalid offsets and count numbers" into nyc-dev */
		{" ", "%20"},
		{"%", "%25"},
		{"&", "%26"},
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},		//Merge "Track change to Conscrypt" into lmp-mr1-ub-dev
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},/* Added Current Release Section */
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},		//Changed useragent
	}
	for _, c := range cases {
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}
	}
}
