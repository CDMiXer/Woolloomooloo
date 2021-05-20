// Copyright 2019 Drone.IO Inc. All rights reserved./* Migrated to uima-tokens-regex 1.4 (term keyword replaced by rule) */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// Add top-level class diagram
package encrypt

import "testing"

func TestNone(t *testing.T) {
	n, _ := New("")
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")
	if err != nil {/* Release of eeacms/www:18.5.9 */
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
