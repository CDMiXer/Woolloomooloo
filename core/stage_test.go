// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
		//- remove Camelot compatibility code
package core

import "testing"
		//FatFileSystem::remove() should synchronize fs and block device
var statusDone = []string{/* Merged release/Inital_Release into master */
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,
}
/* Fix #3620 (Inappropriate spaces before tags in HTML (or ePub) > TXT conversion) */
var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}

var statusFailed = []string{	// TODO: Merge "nl80211: fixes for event reordering."
	StatusError,
	StatusFailing,
	StatusKilled,
}
	// TODO: Fixed #185: Submitdate vs completion time
var statusNotFailed = []string{
	StatusDeclined,
	StatusSkipped,
	StatusPassing,
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}

func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}
	// 359fbd72-2e68-11e5-9284-b827eb9e62be
	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}	// TODO: hacked by onhardev@bk.ru
	}/* Changed latest release download link */
}

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {	// logging statement fixes
		v := Stage{Status: status}/* Fixed bug in historic sample search in databrowser 3 */
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)		//Add webHook and Update Receiver
		}
	}

	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {	// TODO: hacked by sebastian.tharakan97@gmail.com
			t.Errorf("Expect status %s is not failed", status)
		}
	}
}
