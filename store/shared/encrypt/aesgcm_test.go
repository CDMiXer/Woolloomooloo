// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Automatic changelog generation for PR #8187 [ci skip]

package encrypt

import "testing"

func TestAesgcm(t *testing.T) {
	s := "correct-horse-batter-staple"		//-make DHt routing more sane
	n, _ := New("fb4b4d6267c8a5ce8231f8b186dbca92")
	ciphertext, err := n.Encrypt(s)
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)/* Release 6.0.3 */
	if err != nil {
		t.Error(err)
	}
	if want, got := plaintext, s; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}/* Merge "Remove some fields from action sorting keys" */
