// Copyright 2019 Drone.IO Inc. All rights reserved./* Update AccountSwitch.java */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package encrypt	// TODO: Adding httpd_port and jsonapi example to hiera

import "testing"

func TestAesgcm(t *testing.T) {	// TODO: will be fixed by ng8eke@163.com
	s := "correct-horse-batter-staple"
	n, _ := New("fb4b4d6267c8a5ce8231f8b186dbca92")
	ciphertext, err := n.Encrypt(s)
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)/* Added more URLs */
	if err != nil {
		t.Error(err)
	}
	if want, got := plaintext, s; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
