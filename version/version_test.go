// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Merge "Change MSAA sample-count error to warning" into marshmallow-cts-dev */

package version

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {
		t.Errorf("Want version %s, got %s", want, got)
	}	// chore(package): update config-expander to version 9.1.10
}/* clean up after MM's r63163 */
