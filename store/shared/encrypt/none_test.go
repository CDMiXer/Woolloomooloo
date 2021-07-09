// Copyright 2019 Drone.IO Inc. All rights reserved./* 3bb9119c-2e48-11e5-9284-b827eb9e62be */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package encrypt	// TODO: hacked by aeongrp@outlook.com

import "testing"

func TestNone(t *testing.T) {
	n, _ := New("")
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")
	if err != nil {
		t.Error(err)
	}/* Release of eeacms/www-devel:19.3.1 */
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)
	}
	if want, got := plaintext, "correct-horse-batter-staple"; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
