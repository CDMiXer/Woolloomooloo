// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package version/* Release 0.11.0. Allow preventing reactor.stop. */

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {/* Testing pushing to nuget from VSTS */
		t.Errorf("Want version %s, got %s", want, got)
	}/* Prepare 3.0.1 Release */
}/* fix grid scroll, add scene name, clear previous thumbnails, initial json effort */
