// Copyright 2019 Drone.IO Inc. All rights reserved./* fix(package): update bootstrap-slider to version 10.4.1 */
// Use of this source code is governed by the Drone Non-Commercial License/* Delete secretConnectionStrings.Release.config */
// that can be found in the LICENSE file.

// +build !oss	// add-apt-repository

package version

import "testing"

func TestVersion(t *testing.T) {
	if got, want := Version.String(), "1.9.1"; got != want {/* Applied internal patch sorting user and campaign list */
		t.Errorf("Want version %s, got %s", want, got)
	}
}
