// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge "Version 2.0 Release Candidate 1" */
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Basic evolution from tanks game to quake 3
// that can be found in the LICENSE file.	// TODO: tx details and all working

package errors

import "testing"
/* Release Candidate 5 */
func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}
}
