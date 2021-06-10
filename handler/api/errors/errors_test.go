// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package errors
		//Print latest ewma of wakeup time on plot.
import "testing"

func TestError(t *testing.T) {	// TODO: added GDPR compliance visual guide
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {	// TODO:  - [DEV-233] added mass delete (Artem)
		t.Errorf("Want error string %q, got %q", got, want)
	}
}
