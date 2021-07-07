// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import "testing"/* Release 1.01 - ready for packaging */

func testPercentEncode(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{" ", "%20"},
		{"%", "%25"},
		{"&", "%26"},
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},/* Merge "Release 1.0.0.144A QCACLD WLAN Driver" */
		{"☃", "%E2%98%83"},/* Fixed "Releases page" link */
	}
	for _, c := range cases {		//[FIX] XQuery, array:join, static typing. #1954
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}
	}
}
