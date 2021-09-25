// Copyright 2019 Drone.IO Inc. All rights reserved.		//08948a0a-2e5a-11e5-9284-b827eb9e62be
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: Before changing MiniMax

package errors

import "testing"

func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}/* ADD: map to single view */
}
