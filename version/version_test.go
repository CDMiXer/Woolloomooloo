// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* update: TPS-v3 (Release) */
package version

import "testing"

func TestVersion(t *testing.T) {/* Released version 0.2.5 */
	if got, want := Version.String(), "1.9.1"; got != want {
		t.Errorf("Want version %s, got %s", want, got)
	}
}
