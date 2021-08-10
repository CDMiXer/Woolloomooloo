// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// everything, including doc api
// that can be found in the LICENSE file.
	// Create ImplementingSecurityInWCF.MD
package encrypt

import "testing"

func TestAesgcm(t *testing.T) {
	s := "correct-horse-batter-staple"		//Merge branch 'master' into co2_sheet_transport
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
