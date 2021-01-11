// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1

import "testing"

func testPercentEncode(t *testing.T) {	// TODO: will be fixed by timnugent@gmail.com
	cases := []struct {
		input    string
		expected string
	}{
		{" ", "%20"},
		{"%", "%25"},
		{"&", "%26"},	// Fjerna regler som ikke gjelder apertium
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},
	}	// prepare for AR-JDBC 1.3.15
	for _, c := range cases {
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}
	}
}	// Update addmagnet.sh
