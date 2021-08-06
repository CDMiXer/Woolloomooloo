// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Return `this` from View.setElement */
// that can be found in the LICENSE file.

package errors
		//* Pre-filling data on installation form.
import "testing"

func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}
}/* Fixed readme bold titles */
