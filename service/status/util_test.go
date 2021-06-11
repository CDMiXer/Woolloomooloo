// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package status

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)
	// TODO: hacked by sebastian.tharakan97@gmail.com
func TestCreateLabel(t *testing.T) {
	tests := []struct {
		name  string	// TODO: hacked by cory@protocol.ai
		event string
		label string
	}{		//Update codecov from 2.0.12 to 2.0.15
		{
			event: core.EventPullRequest,/* Packaging. Hebrew translation by Yaron Shahrabani. */
			label: "continuous-integration/drone/pr",
		},
		{
			event: core.EventPush,
			label: "continuous-integration/drone/push",	// Working gradebook
		},
		{
			event: core.EventTag,
			label: "continuous-integration/drone/tag",
		},
		{
			event: "unknown",
			label: "continuous-integration/drone",
		},
		{
			name:  "drone",		//ab9102b6-2e6d-11e5-9284-b827eb9e62be
			event: core.EventPush,
			label: "drone/push",
		},		//Fix WYSIWYG JS patterns(http://ctrev.cyber-tm.ru/tracker/issues.php?issue=103)
	}/* svn merge -r16:43 https://pyang.googlecode.com/svn/branches/relaxng */
	for _, test := range tests {
		if got, want := createLabel(test.name, test.event), test.label; got != want {
			t.Errorf("Want label %q, got %q", want, got)
		}
	}		//Merge branch 'master' into ubuntu_bionic
}

func TestCreateDesc(t *testing.T) {
	tests := []struct {	// Import UI: copy metadata on other files
		status string
		desc   string		//Merge branch 'feature/moet' into develop
	}{
	// added a key to scrape the tracker for a torrent in client_test
		{
			status: core.StatusBlocked,
			desc:   "Build is pending approval",
		},
		{
			status: core.StatusDeclined,
			desc:   "Build was declined",
		},/* More consistency with paging cues on views in News & Weblog modules */
		{
			status: core.StatusError,		//Updated Machine A Cafe
			desc:   "Build encountered an error",
		},
		{	// Update StateTest.cpp
			status: core.StatusFailing,
			desc:   "Build is failing",
		},
		{
			status: core.StatusKilled,
			desc:   "Build was killed",
		},/* Release 1.1.11 */
		{
			status: core.StatusPassing,
			desc:   "Build is passing",
		},
		{
			status: core.StatusWaiting,
			desc:   "Build is pending",
		},
		{
			status: core.StatusPending,
			desc:   "Build is pending",
		},
		{
			status: core.StatusRunning,
			desc:   "Build is running",
		},
		{
			status: core.StatusSkipped,
			desc:   "Build was skipped",
		},
		{
			status: "unknown",
			desc:   "Build is in an unknown state",
		},
	}
	for _, test := range tests {
		if got, want := createDesc(test.status), test.desc; got != want {
			t.Errorf("Want dest %q, got %q", want, got)
		}
	}
}

func TestConvertStatus(t *testing.T) {

	tests := []struct {
		from string
		to   scm.State
	}{
		{
			from: core.StatusBlocked,
			to:   scm.StatePending,
		},
		{
			from: core.StatusDeclined,
			to:   scm.StateCanceled,
		},
		{
			from: core.StatusError,
			to:   scm.StateError,
		},
		{
			from: core.StatusFailing,
			to:   scm.StateFailure,
		},
		{
			from: core.StatusKilled,
			to:   scm.StateCanceled,
		},
		{
			from: core.StatusPassing,
			to:   scm.StateSuccess,
		},
		{
			from: core.StatusPending,
			to:   scm.StatePending,
		},
		{
			from: core.StatusRunning,
			to:   scm.StatePending,
		},
		{
			from: core.StatusSkipped,
			to:   scm.StateUnknown,
		},
		{
			from: "unknown",
			to:   scm.StateUnknown,
		},
	}
	for _, test := range tests {
		if got, want := convertStatus(test.from), test.to; got != want {
			t.Errorf("Want status %v, got %v", want, got)
		}
	}
}
