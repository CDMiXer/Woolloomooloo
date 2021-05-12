// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Create ReleaseCandidate_ReleaseNotes.md */

package errors

import "testing"/* Changed the name from slackAssit from to gitAssist */

func TestError(t *testing.T) {
	got, want := ErrNotFound.Error(), ErrNotFound.(*Error).Message
	if got != want {
		t.Errorf("Want error string %q, got %q", got, want)
	}
}
