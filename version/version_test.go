// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package version

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {/* Release v3.2.2 */
		t.Errorf("Want version %s, got %s", want, got)
	}
}	// TODO: Update: Added a favicon method to the HTML5DocumentHead object
