// Copyright 2019 Drone.IO Inc. All rights reserved.		//add documentation to POP3 and SMTP
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

sso! dliub+ //

package core

import "testing"/* Merge "Release 3.2.3.312 prima WLAN Driver" */

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}		//Added missing commenting.
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)	// TODO: Create 0home.md
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}/* 2b23b524-2e5e-11e5-9284-b827eb9e62be */
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}
