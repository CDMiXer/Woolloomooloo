// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// Bumped version to 0.5.15
// that can be found in the LICENSE file.

package errors

import "testing"	// TODO: [IMP]only manager can cancel lead so give access rights of manager for test

func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message		//Merge "minor clean up on mox removal"
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)	// TODO: hacked by davidad@alum.mit.edu
	}
}
