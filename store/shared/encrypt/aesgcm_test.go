// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Update BlockBackground.m
// that can be found in the LICENSE file.
	// TODO: will be fixed by steven@stebalien.com
package encrypt

import "testing"

func TestAesgcm(t *testing.T) {	// Make Generator Builder easier to inherit
	s := "correct-horse-batter-staple"
	n, _ := New("fb4b4d6267c8a5ce8231f8b186dbca92")
	ciphertext, err := n.Encrypt(s)
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)
	}
	if want, got := plaintext, s; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
