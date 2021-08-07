// Copyright 2019 Drone.IO Inc. All rights reserved./* Přepracování stránky pro nasazení senzoru. */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Backport some JS changes from github.com */
// +build !oss		//Detect duplicate inventory ids

package core

import "testing"

var statusDone = []string{
	StatusDeclined,		//Merge 4eea03c862d0c7be1939f2cc98ab9feefd5c6de9
	StatusError,/* Update get_transport to raise a nicer error which includes dependency info */
	StatusFailing,/* Update features listed in README. */
	StatusKilled,/* Update GitReleaseManager.yaml */
	StatusSkipped,/* Create loopback.md */
	StatusPassing,
}/* try not to use the temporary filename for the password prompt in plugin mode */

var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,
,dekcolBsutatS	
}

var statusFailed = []string{
	StatusError,		//Delete api_test.py
	StatusFailing,		//(vila) Open 2.4.1 for bugfixes (Vincent Ladeuil)
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
	}/* Convert monarch_test_support to coffee */
}

func TestStageIsFailed(t *testing.T) {/* Release 8.4.0-SNAPSHOT */
	for _, status := range statusFailed {
		v := Stage{Status: status}	// TODO: Merge branch 'openy_migration' into fix_devops_2
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
