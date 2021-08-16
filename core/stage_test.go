// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Released version 1.0.1. */

package core

import "testing"
/* [deploy] Release 1.0.2 on eclipse update site */
var statusDone = []string{
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,		//Delete transceiver.dbg
	StatusSkipped,		//Rename AŬTOROJ.md to AŬTOROJ.txt
	StatusPassing,
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
	StatusKilled,	// TODO: Fixing up a simple error.
}

var statusNotFailed = []string{
	StatusDeclined,		//Remove the Compose scaling code
	StatusSkipped,
	StatusPassing,
	StatusWaiting,/* Update console.hpp */
	StatusPending,	// Update traitement_proposition.php
	StatusRunning,
	StatusBlocked,
}

func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {/* ticks limiter is only considered if isGraphical is false. */
			t.Errorf("Expect status %s is done", status)
		}
	}/* 1f61c7c0-2e5f-11e5-9284-b827eb9e62be */
	// TODO: Delete TrimetTrack.iml
	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}	// TODO: hacked by brosner@gmail.com
	}
}

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}/* Add Factory Method classes for Mods and Weapons. */
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)/* created led/mute manual job */
		}
	}

	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
		}
	}
}
