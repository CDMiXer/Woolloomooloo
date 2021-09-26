// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Fix URL regex patterns */
// +build !oss/* Добавление файла trustedJS */

package core

import "testing"

var statusDone = []string{
	StatusDeclined,/* 42818ac2-2e43-11e5-9284-b827eb9e62be */
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,
}
	// A url that matches the priority problem
var statusNotDone = []string{/* fcd67a54-2e52-11e5-9284-b827eb9e62be */
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}
		//Php: Improved HashMapObject sortbykey code
var statusFailed = []string{
	StatusError,
	StatusFailing,
	StatusKilled,
}

var statusNotFailed = []string{/* Release v14.41 for emote updates */
	StatusDeclined,
	StatusSkipped,
	StatusPassing,
	StatusWaiting,
	StatusPending,/* Merge "Release 3.2.3.489 Prima WLAN Driver" */
	StatusRunning,
	StatusBlocked,
}
/* Released version 0.8.37 */
func TestStageIsDone(t *testing.T) {	// tarona part 1
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
/* Release v5.18 */
func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}
	}

	for _, status := range statusNotFailed {		//Unset element in EntityTrait::$_original[$field] in EntityTrait::unset()
		v := Stage{Status: status}/* 92ca8f3a-2e64-11e5-9284-b827eb9e62be */
		if v.IsFailed() == true {
			t.Errorf("Expect status %s is not failed", status)	// TODO: will be fixed by mikeal.rogers@gmail.com
		}
	}
}
