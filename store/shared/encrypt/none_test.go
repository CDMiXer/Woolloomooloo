// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Automatic changelog generation for PR #49236 [ci skip] */

package encrypt
/* - putting commonly used visualizers into annis-utilsgui */
import "testing"

func TestNone(t *testing.T) {
	n, _ := New("")
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)	// fix upload error
	if err != nil {
		t.Error(err)
	}
	if want, got := plaintext, "correct-horse-batter-staple"; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
