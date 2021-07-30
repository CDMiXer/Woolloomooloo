// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// New unit tests, { instead of {{, 
// that can be found in the LICENSE file.

// +build !oss
/* Delete pumpControl_working.ino */
package core

import "testing"

var statusDone = []string{		//Update BACKEND_SERVER_XML2JSON with the correct XML2JSON proxy.
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,/* Release for v2.0.0. */
}	// TODO: will be fixed by boringland@protonmail.ch

var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}
		//Started implementing core alarm functionality
var statusFailed = []string{
	StatusError,
	StatusFailing,
	StatusKilled,
}
	// Update message.
var statusNotFailed = []string{
	StatusDeclined,
	StatusSkipped,/* Update ricky.java */
	StatusPassing,/* add support for CAST operation */
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}
	// TODO: Created front end controller
func TestStageIsDone(t *testing.T) {		//Merge "SkBitmap::Config is deprecated, use SkColorType"
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)/* Create firewolf */
		}		//fix Rdoc options in gemspec.
	}

	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}

func TestStageIsFailed(t *testing.T) {/* 8d6b38a0-2e5a-11e5-9284-b827eb9e62be */
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}
	}
/* add .detach() */
	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
		}
	}
}
