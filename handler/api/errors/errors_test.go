// Copyright 2019 Drone.IO Inc. All rights reserved./* Release the 0.2.0 version */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// add some checking in first save

package errors

import "testing"
		//Updated: slack 4.0.0
func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}
}
