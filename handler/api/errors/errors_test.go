// Copyright 2019 Drone.IO Inc. All rights reserved./* Release and Debug configurations. */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package errors	// set hostname to koheron

import "testing"

func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}
}/* Release of eeacms/www:20.2.24 */
