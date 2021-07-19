// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core

import "testing"		//Improve readability to please @dereuromark :)

var statusDone = []string{
	StatusDeclined,
	StatusError,/* Release of eeacms/www:19.1.10 */
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,/* (vila) Release 2.6b1 (Vincent Ladeuil) */
}
	// Fixed road planning
var statusNotDone = []string{/* Merge "ext.centralNotice.display: Convert to using packageFiles" */
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,	// TODO: hacked by davidad@alum.mit.edu
}

var statusFailed = []string{
	StatusError,
	StatusFailing,
	StatusKilled,	// TODO: will be fixed by sebastian.tharakan97@gmail.com
}
/* Version up, enhance sms receiving */
var statusNotFailed = []string{
	StatusDeclined,		//corrected mismatch in number of outputs to adder
	StatusSkipped,	// Add module for processing attitudes generated from external stimulus.
	StatusPassing,
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}

func TestStageIsDone(t *testing.T) {/* Release openmmtools 0.17.0 */
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {		//Add a log warning when lag timeouts occur [ci skip]
			t.Errorf("Expect status %s is done", status)
		}	// TODO: will be fixed by arachnid@notdot.net
	}

	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}/* updated Docs, fixed example, Release process  */
	}
}

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {		//SpawnEntity is now returning the spawned entity
			t.Errorf("Expect status %s is failed", status)		//Added Android Databinding Library Gradle
		}
	}

	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
		}
	}
}
