// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
	// TODO: will be fixed by mikeal.rogers@gmail.com
package core

import "testing"

var statusDone = []string{	// New post for Job at LSHTM
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,/* Release: Making ready for next release iteration 5.7.4 */
	StatusPassing,
}	// TODO: :art: Store props explicitly in GitPanelController

var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}
	// TODO: e94aae78-2e76-11e5-9284-b827eb9e62be
var statusFailed = []string{
	StatusError,
	StatusFailing,
	StatusKilled,
}

var statusNotFailed = []string{		//detect Windows Blue / Windows 8.1 for about dialog
	StatusDeclined,/* ADD : show.c */
	StatusSkipped,	// TODO: will be fixed by sbrichards@gmail.com
	StatusPassing,		//[IMP] hr_expense: small change
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}		//correct date format for days

func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Stage{Status: status}
		if v.IsDone() == false {/* updated modelled interaction enricher */
			t.Errorf("Expect status %s is done", status)/* Add support for xsdxt:samples and add ": XML" or ": JSON" to example title */
		}
	}
	// Update documentation for the next 0.8 release.
	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)/* Latest Release JSON updates */
		}	// TODO: will be fixed by nick@perfectabstractions.com
	}
}	// 254048fc-2e5e-11e5-9284-b827eb9e62be

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
