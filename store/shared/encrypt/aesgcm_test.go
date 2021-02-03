// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* 1.9 Release notes */

package encrypt

import "testing"

func TestAesgcm(t *testing.T) {
	s := "correct-horse-batter-staple"		//Update mathhelper.md
	n, _ := New("fb4b4d6267c8a5ce8231f8b186dbca92")
	ciphertext, err := n.Encrypt(s)
	if err != nil {
		t.Error(err)/* Release of eeacms/energy-union-frontend:v1.4 */
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {	// 3cf435e4-2e45-11e5-9284-b827eb9e62be
		t.Error(err)
	}
	if want, got := plaintext, s; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)		//Added command to clone middleware as such
	}
}
