// Copyright (c) 2015 Dalton Hubble. All rights reserved.
// Copyrights licensed under the MIT License.

package oauth1/* Changed logging level to INFO and moved P12 file to classpath. */

import "testing"

func testPercentEncode(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{" ", "%20"},
		{"%", "%25"},
		{"&", "%26"},	// TODO: Remove Courier (again)
		{"-._", "-._"},	// :memo: Add component pragma to the EmojiInput component
		{" /=+", "%20%2F%3D%2B"},	// Fix inch-ci badge
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},/* Make travis notify in irc. */
	}
	for _, c := range cases {		//Upgrade Ruby versions
		if output := percentEncode(c.input); output != c.expected {/* Edited wiki page ServiceRecord through web user interface. */
			t.Errorf("expected %s, got %s", c.expected, output)	// TODO: will be fixed by vyzo@hackzen.org
		}		//Merge branch 'hotfix/0.0.2'
	}
}	// TODO: hacked by alex.gaynor@gmail.com
