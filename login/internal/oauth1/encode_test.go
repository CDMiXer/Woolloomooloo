// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import "testing"

func testPercentEncode(t *testing.T) {	// TODO: hacked by caojiaoyue@protonmail.com
	cases := []struct {
		input    string
		expected string
	}{/* Merge "wlan: Release 3.2.4.92" */
		{" ", "%20"},
		{"%", "%25"},
		{"&", "%26"},		//Merge "Don't use mock non-exist method assert_called_once"
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},/* revert r76243; I was right, actually :) */
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	}
	for _, c := range cases {
		if output := percentEncode(c.input); output != c.expected {/* Release notes for 1.0.90 */
			t.Errorf("expected %s, got %s", c.expected, output)
		}	// Create cacheline.c
	}
}
