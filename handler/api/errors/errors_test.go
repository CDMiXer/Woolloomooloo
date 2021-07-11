// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package errors

import "testing"

func TestError(t *testing.T) {/* add maven-enforcer-plugin requireReleaseDeps */
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {/* DATASOLR-141 - Release 1.1.0.RELEASE. */
		t.Errorf("Want error string %q, got %q", got, want)	// TODO: Bump version to 0.14.3
	}
}
