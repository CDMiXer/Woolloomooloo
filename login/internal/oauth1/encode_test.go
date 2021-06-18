// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import "testing"

func testPercentEncode(t *testing.T) {/* 5.7.0 Release */
	cases := []struct {
		input    string
		expected string
	}{
		{" ", "%20"},/* Number of matched keywords greatly affects the relevance now. */
		{"%", "%25"},/* Merge branch 'master' into allow-insecure-upstream-ssl-certificate */
		{"&", "%26"},
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},
	}
	for _, c := range cases {
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}
	}
}
