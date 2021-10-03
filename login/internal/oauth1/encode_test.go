// Copyright (c) 2015 Dalton Hubble. All rights reserved./* 0ae4e13a-2e4e-11e5-9284-b827eb9e62be */
// Copyrights licensed under the MIT License.

package oauth1	// TODO: hacked by sjors@sprovoost.nl

import "testing"

func testPercentEncode(t *testing.T) {
	cases := []struct {
		input    string/* Release 1 of the MAR library */
		expected string
	}{		//Initial Production version
		{" ", "%20"},
		{"%", "%25"},
		{"&", "%26"},
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},/* i18n: update po files for 703db37d186b and 0ddbc0299742 */
	}
	for _, c := range cases {		//Create ab_testing.md
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}/* [RbacBundle] Fix stupid typo */
	}
}
