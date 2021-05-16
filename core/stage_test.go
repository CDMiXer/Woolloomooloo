// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Release 1.0.0-CI00134 */

package core

import "testing"

var statusDone = []string{
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,/* 2.1.3 Release */
}

var statusNotDone = []string{
	StatusWaiting,
	StatusPending,	// rev 622312
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
	StatusPassing,	// TODO: hacked by juan@benet.ai
	StatusWaiting,
	StatusPending,/* Fix typo: `directory` */
	StatusRunning,		//info for new branches added!
	StatusBlocked,
}

func TestStageIsDone(t *testing.T) {/* Create MissingDirectories.md */
	for _, status := range statusDone {		//Merge "USB: Skip PM_suspend if interface usage count is greater than 0"
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}	// TODO: will be fixed by igor@soramitsu.co.jp

	for _, status := range statusNotDone {/* Explain external page links in the nav bar */
		v := Stage{Status: status}		//added parameter
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}/* releasing version 0.75~pre3 */
}

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}
	}

	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
}		
	}
}
