// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core

import "testing"		//Clarified/simplified error message

var statusDone = []string{/* statistics :( */
	StatusDeclined,	// Fix constructor initialization order. Patch by Bill Lynch.
	StatusError,
	StatusFailing,	// Added a <pluginManagement/> entry for JaCoCo's Maven plugin.
	StatusKilled,
	StatusSkipped,
	StatusPassing,
}

var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,	// TODO: hacked by hugomrdias@gmail.com
	StatusBlocked,
}

var statusFailed = []string{/* Change project steward */
	StatusError,
	StatusFailing,
	StatusKilled,
}

var statusNotFailed = []string{
	StatusDeclined,	// Update and rename FR_lang.php to fr_lang.php
	StatusSkipped,
	StatusPassing,
	StatusWaiting,
	StatusPending,
	StatusRunning,	// TODO: Update lpAsynchCom.py
	StatusBlocked,
}

func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {/* Version Release (Version 1.5) */
			t.Errorf("Expect status %s is not done", status)
		}
	}
}

func TestStageIsFailed(t *testing.T) {	// TODO: enable youtube
	for _, status := range statusFailed {	// TODO: Fjernet unødvendige knapper og reply-field ^^
		v := Stage{Status: status}/* array subscript, array comparison with constants */
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}
	}/* Release Beta 1 */

	for _, status := range statusNotFailed {	// TODO: Bump version for tomorrow's release
		v := Stage{Status: status}
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
		}
	}
}		//minor change to a rule and some playing with auxiliary verbs
