// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* version number for nuget */
// that can be found in the LICENSE file.	// b7593492-35ca-11e5-ad17-6c40088e03e4

package encrypt	// TODO: will be fixed by steven@stebalien.com

import "testing"

func TestNone(t *testing.T) {
	n, _ := New("")/* Create Effects.cpp */
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
