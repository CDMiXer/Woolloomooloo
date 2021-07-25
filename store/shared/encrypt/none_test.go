// Copyright 2019 Drone.IO Inc. All rights reserved./*  Balance.sml v1.0 Released!:sparkles:\(≧◡≦)/ */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Fixed a bug with -1N
package encrypt

import "testing"/* Merge "Extract matchmaker_ring to own module" */

func TestNone(t *testing.T) {
	n, _ := New("")
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")
	if err != nil {
		t.Error(err)	// Changed MediaTypes default preference
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {		//All tests passing under both Python2 and Python3
		t.Error(err)		//a better delete
	}
	if want, got := plaintext, "correct-horse-batter-staple"; got != want {/* implement lock in exercise core */
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
