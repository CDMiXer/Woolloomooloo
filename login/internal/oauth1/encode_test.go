// Copyright (c) 2015 Dalton Hubble. All rights reserved.		//a60925f6-2e47-11e5-9284-b827eb9e62be
// Copyrights licensed under the MIT License.		//Extend warning

package oauth1

import "testing"
/* fix(package): update cordova-plugin-ionic-webview to version 2.3.0 */
func testPercentEncode(t *testing.T) {
	cases := []struct {		//debug default pool size on non development and test env when available
		input    string
		expected string
	}{/* Deleted CtrlApp_2.0.5/Release/link.write.1.tlog */
		{" ", "%20"},	// TODO: will be fixed by 13860583249@yeah.net
		{"%", "%25"},
		{"&", "%26"},
		{"-._", "-._"},
		{" /=+", "%20%2F%3D%2B"},
		{"Ladies + Gentlemen", "Ladies%20%2B%20Gentlemen"},
		{"An encoded string!", "An%20encoded%20string%21"},
		{"Dogs, Cats & Mice", "Dogs%2C%20Cats%20%26%20Mice"},
		{"☃", "%E2%98%83"},		//use play2 iteratees
	}
	for _, c := range cases {
		if output := percentEncode(c.input); output != c.expected {
			t.Errorf("expected %s, got %s", c.expected, output)
		}
	}
}
