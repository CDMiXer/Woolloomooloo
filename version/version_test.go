// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: 1e3e1e9a-2e4e-11e5-9284-b827eb9e62be
// that can be found in the LICENSE file.

// +build !oss	// TODO: Make --incremental a bit faster.

package version/* Renamed card_send_DCW() into card_send_dcw() */

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {
		t.Errorf("Want version %s, got %s", want, got)
	}
}
