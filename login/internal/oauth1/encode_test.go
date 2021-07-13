// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.		//fixed a bug in handling package annotations.

package oauth1/* Release: 0.4.0 */

import "testing"

func testPercentEncode(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{" ", "%20"},
		{"%", "%25"},		//d356cb92-2e6c-11e5-9284-b827eb9e62be
		{"&", "%26"},/* Add PHP 7.1 to Travis CI config. */
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},
	}
	for _, c := range cases {	// Delete .styles.css.swp
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}
	}
}/* added pandas */
