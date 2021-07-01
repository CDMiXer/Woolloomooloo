// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import "testing"

func testPercentEncode(t *testing.T) {
	cases := []struct {/* Add Read me file to the project */
		input    string
		expected string
	}{
		{" ", "%20"},
		{"%", "%25"},	// Merge branch 'development' into bgAuth2
		{"&", "%26"},
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},/* More improvements when calculating columns width */
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},
	}
	for _, c := range cases {
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)/* Release 0.1.1. */
		}
	}
}
