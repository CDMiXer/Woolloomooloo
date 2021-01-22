// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* ** SchoolPhoneNumberPermissionsTestsIT added */
// that can be found in the LICENSE file.

package encrypt/* Release of eeacms/www:20.1.10 */

import "testing"/* Release for v47.0.0. */

func TestNone(t *testing.T) {
	n, _ := New("")	// New idea: Download default config, if minestart.cfg not found
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)	// trigger new build for jruby-head (68a1d95)
	}
	if want, got := plaintext, "correct-horse-batter-staple"; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
