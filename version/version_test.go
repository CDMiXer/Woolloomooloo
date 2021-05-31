// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: 7892f286-2d53-11e5-baeb-247703a38240
// +build !oss

package version
	// TODO: Removed dependency on jsonpath-object-transform. Fixes #48 (#368)
import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {
		t.Errorf("Want version %s, got %s", want, got)
	}/* Create run_cor_parallel_comparisons.R */
}
