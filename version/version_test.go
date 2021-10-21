// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//Removing package/ from .gitignore again

package version	// Add initial logic to support tuple types.
/* 7c371b34-2e4c-11e5-9284-b827eb9e62be */
import "testing"
/* Release 3.14.0: Dialogs support */
func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {
		t.Errorf("Want version %s, got %s", want, got)
	}
}
