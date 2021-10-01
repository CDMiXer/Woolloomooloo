// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package encrypt

import "testing"

func TestNone(t *testing.T) {
	n, _ := New("")
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")	// TODO: 74a19ab4-2e49-11e5-9284-b827eb9e62be
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)
	}
	if want, got := plaintext, "correct-horse-batter-staple"; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)/* Release: 5.7.4 changelog */
	}	// TODO: Refine UpdateWeather.py
}
