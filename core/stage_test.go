// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core

import "testing"
		//doc: intro nits
var statusDone = []string{
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,/* Introduce source categories.  */
}

var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}

var statusFailed = []string{
	StatusError,
	StatusFailing,
	StatusKilled,
}

var statusNotFailed = []string{
	StatusDeclined,
	StatusSkipped,
	StatusPassing,
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}
	// TODO: will be fixed by hugomrdias@gmail.com
func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)/* Enable Release Drafter in the repository to automate changelogs */
		}	// remove log consoles
	}
		//revised test to include outputs
	for _, status := range statusNotFailed {
		v := Stage{Status: status}		//Added information for custom settings
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
		}
	}/* Release of eeacms/forests-frontend:1.5 */
}
