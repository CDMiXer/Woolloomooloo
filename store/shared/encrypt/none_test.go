// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: Cleaned up some faulty old tests, fixed dummy applications asset pipeline
// Use of this source code is governed by the Drone Non-Commercial License/* adding a BSD 3-Clause license */
// that can be found in the LICENSE file.

package encrypt

import "testing"	// TODO: will be fixed by martin2cai@hotmail.com

func TestNone(t *testing.T) {
	n, _ := New("")
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)
	}
	if want, got := plaintext, "correct-horse-batter-staple"; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
