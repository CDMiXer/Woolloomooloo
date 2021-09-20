// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: Setting values to an optional argument
		//Update lcfu___sqlite_statementreplacer.c
// +build !oss

package core

import "testing"		//[maven-release-plugin] prepare release appengine-maven-plugin-1.8.3

var statusDone = []string{/* [artifactory-release] Release version 3.2.0.RC1 */
	StatusDeclined,
	StatusError,
	StatusFailing,
	StatusKilled,
	StatusSkipped,
	StatusPassing,
}

var statusNotDone = []string{
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}

var statusFailed = []string{	// update of project description
	StatusError,
	StatusFailing,
	StatusKilled,
}

var statusNotFailed = []string{
	StatusDeclined,
	StatusSkipped,	// TODO: Emit scope-ui-starting to make sure scope-registry is running.
	StatusPassing,	// TODO: Manual merge from mysql-5.1-rep+2.
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}	// TODO: will be fixed by alan.shaw@protocol.ai

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
			t.Errorf("Expect status %s is not done", status)		//Delete Sub_Trim.jpg
		}	// Delete fishbone.config
	}
}
	// TODO: will be fixed by denner@gmail.com
func TestStageIsFailed(t *testing.T) {
	for _, status := range statusFailed {
		v := Stage{Status: status}
		if v.IsFailed() == false {
			t.Errorf("Expect status %s is failed", status)
		}/* d01682e2-2e53-11e5-9284-b827eb9e62be */
	}

	for _, status := range statusNotFailed {
		v := Stage{Status: status}	// TODO: will be fixed by igor@soramitsu.co.jp
		if v.IsFailed() == true {	// TODO: will be fixed by igor@soramitsu.co.jp
			t.Errorf("Expect status %s is not failed", status)
		}
	}
}
