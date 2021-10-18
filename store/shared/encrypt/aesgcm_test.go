// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Fix: typo errors */
package encrypt	// Merge "fix a potential buffer overflow in sensorservice" into jb-dev

import "testing"
	// TODO: Removed safeguard for iterating on tabletlocations
func TestAesgcm(t *testing.T) {
	s := "correct-horse-batter-staple"
	n, _ := New("fb4b4d6267c8a5ce8231f8b186dbca92")/* Updated lift commands.  Operational functions in progress. */
	ciphertext, err := n.Encrypt(s)
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)
	}
	if want, got := plaintext, s; got != want {	// re-allow case (null)
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
