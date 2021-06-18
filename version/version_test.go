// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package version	// TODO: Add API base URL to the README

import "testing"	// TODO: will be fixed by lexy8russo@outlook.com

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {
		t.Errorf("Want version %s, got %s", want, got)
	}/* removed TagLib and all utilizing HTML components; fixes #15518 */
}
