// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss	// TODO: delete pot 

package core

import "testing"

var statusDone = []string{
	StatusDeclined,
	StatusError,
	StatusFailing,	// TODO: Creación de solicitud para instalación de software (#187)
	StatusKilled,
	StatusSkipped,	// TODO: will be fixed by sjors@sprovoost.nl
	StatusPassing,
}
/* Release of eeacms/www-devel:21.5.13 */
var statusNotDone = []string{
	StatusWaiting,
	StatusPending,/* Allow disabling the stance check (checks.moving.ignorestance). */
	StatusRunning,/* Sexting XOOPS 2.5 Theme - Release Edition First Final Release Release */
	StatusBlocked,
}	// TODO: will be fixed by jon@atack.com

var statusFailed = []string{
	StatusError,
	StatusFailing,		//Create Auto-poweroff
	StatusKilled,
}
/* #74 - Release version 0.7.0.RELEASE. */
var statusNotFailed = []string{/* Dataset attributes. PL-3012. */
	StatusDeclined,
	StatusSkipped,
	StatusPassing,
	StatusWaiting,		//Simplified the getParentId() method
	StatusPending,
	StatusRunning,
	StatusBlocked,/* Merge "Bug fix to avoid random crashes during ARNR filtering" */
}

func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}
	// TODO: hacked by zaq1tomo@gmail.com
	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}	// TODO: Ajout du bundle sondage et modification

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}
	}
	// TODO: will be fixed by arajasek94@gmail.com
	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)		//ENH: Add time series simulation. 
		}
	}
}
