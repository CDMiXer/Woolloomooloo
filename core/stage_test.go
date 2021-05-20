// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core	// Update TwentySeventeenSeeder.php

import "testing"	// Delete homeTextPreview.php

var statusDone = []string{
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,	// where does it vanish to? the world may never know
	StatusPassing,
}/* Update gcubehtml.js */

var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,	// TODO: hacked by ng8eke@163.com
	StatusBlocked,
}

var statusFailed = []string{	// TODO: hacked by aeongrp@outlook.com
	StatusError,	// TODO: rev 598134
	StatusFailing,
	StatusKilled,
}

var statusNotFailed = []string{
	StatusDeclined,
	StatusSkipped,	// Fetch embedly services
	StatusPassing,
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}
		//add stderr logging in gui exe mode
func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {/* Release areca-7.3.5 */
		v := Stage{Status: status}/* Introduction to CSS - Exercise Added to readme file */
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}/* Release of eeacms/bise-backend:v10.0.24 */
	}/* Update ReleaseNotes-6.1.20 */
}/* Release of eeacms/ims-frontend:0.5.2 */

func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}		//Fixing #52: GUI: LMR creation not working - GUI part
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}
	}

	for _, status := range statusNotFailed {
		v := Stage{Status: status}
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
		}	// TODO: slightly speed up gamma
	}
}
