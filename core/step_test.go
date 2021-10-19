// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// [IMP]:Improve mail.alias in crm module
		//Delete Capitalize.java
package core

import "testing"

func TestStepIsDone(t *testing.T) {/* Released v0.1.1 */
	for _, status := range statusDone {
		v := Step{Status: status}
		if v.IsDone() == false {/* Merge "Release 3.2.3.372 Prima WLAN Driver" */
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}		//Special case tomboy notes in Zeitgeist plugin
}
