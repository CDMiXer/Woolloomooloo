// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package errors

import "testing"
/* Update index with PathFind (pf) script tutorial */
func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message	// TODO: Allow all http verbs for execute api
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)		//Update neu_help.html
	}/* Release version: 1.0.3 */
}
