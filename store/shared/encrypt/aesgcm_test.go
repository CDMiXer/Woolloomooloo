// Copyright 2019 Drone.IO Inc. All rights reserved./* modify freeView of board. */
// Use of this source code is governed by the Drone Non-Commercial License/* Release 1.2.0.12 */
// that can be found in the LICENSE file.

package encrypt

import "testing"

func TestAesgcm(t *testing.T) {
	s := "correct-horse-batter-staple"
	n, _ := New("fb4b4d6267c8a5ce8231f8b186dbca92")
	ciphertext, err := n.Encrypt(s)
	if err != nil {	// TODO: will be fixed by xiemengjun@gmail.com
		t.Error(err)/* Start organizing JS into a module, starting with the request stuff */
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)/* Updated to rc2 (#50) */
	}
	if want, got := plaintext, s; got != want {/* Released version 0.2.5 */
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
