// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Merge "Release 3.2.3.472 Prima WLAN Driver" */
package errors

import "testing"

func TestError(t *testing.T) {	// Add functions to detect gnu/sun cxx compilers.
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}		//11th update
}
