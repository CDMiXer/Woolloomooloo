// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Initial Release for APEX 4.2.x */
package core/* improve fwd man and connection manager */

import "testing"

var statusDone = []string{
	StatusDeclined,
	StatusError,
	StatusFailing,/* Release 0.23.0 */
	StatusKilled,
	StatusSkipped,
	StatusPassing,
}

var statusNotDone = []string{
	StatusWaiting,/* Release 1.0.13 */
	StatusPending,/* Release version 1.8.0 */
	StatusRunning,	// TODO: hacked by mail@bitpshr.net
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
	StatusRunning,/* Fixed opening using IndexMap */
	StatusBlocked,
}

func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}	// added play_door_sound

	for _, status := range statusNotDone {		//Update genericpostlogin.xhtml
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}

func TestStageIsFailed(t *testing.T) {	// TODO: Renaming label for symmetry.
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}
	}		//Chargement unique de la liste des oeuvres

	for _, status := range statusNotFailed {
		v := Stage{Status: status}/* Unbreak JS API by making sure config is initialized correctly */
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)
		}		//Create girls_ziu_platform.php
	}
}
