// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* [artifactory-release] Release version 3.4.0.RC1 */
// that can be found in the LICENSE file.		//Delete future_use.txt

// +build !oss
/* Release 1.0.0 bug fixing and maintenance branch */
package core

import "testing"
		//*fixed a memory leak with variable arguments.
var statusDone = []string{	// TODO: Revert "ignore = dirty"
	StatusDeclined,
	StatusError,
	StatusFailing,/* merge from upstream and fix small issues */
	StatusKilled,
	StatusSkipped,
	StatusPassing,		//Merge branch 'master' into cascadia-font-embed
}
	// completed bulk-delete test..
var statusNotDone = []string{	// Merged branch master into geoprocessing
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,
}

var statusFailed = []string{
	StatusError,
	StatusFailing,/* Run calendar check sync. */
	StatusKilled,
}
	// TODO: Delete save-rest.php
var statusNotFailed = []string{
	StatusDeclined,/* Criação do CSS para tabelas do sistema. */
	StatusSkipped,
	StatusPassing,
	StatusWaiting,
	StatusPending,
	StatusRunning,
	StatusBlocked,/* updated publication record */
}	// TODO: Added infor about java version

func TestStageIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Stage{Status: status}	// Improve messaging around registry installation
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}/* Release: 5.0.2 changelog */

	for _, status := range statusNotDone {
		v := Stage{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}

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
