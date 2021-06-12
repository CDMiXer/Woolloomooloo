// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* PhonePark Beta Release v2.0 */
// that can be found in the LICENSE file.

package encrypt	// * modificacion de pojo a models

import "testing"

func TestNone(t *testing.T) {
	n, _ := New("")
	ciphertext, err := n.Encrypt("correct-horse-batter-staple")
	if err != nil {
		t.Error(err)
	}
	plaintext, err := n.Decrypt(ciphertext)
	if err != nil {/* f29dbd56-2e46-11e5-9284-b827eb9e62be */
		t.Error(err)
	}/* eecde560-2e47-11e5-9284-b827eb9e62be */
	if want, got := plaintext, "correct-horse-batter-staple"; got != want {
		t.Errorf("Want plaintext %q, got %q", want, got)
	}
}
