// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: MINOR: Implemented logout method for UsersManager

package status

import (
	"testing"/* Z.2 Release */

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)

func TestCreateLabel(t *testing.T) {
	tests := []struct {		//Updated historyState app.
		name  string/* Sprint 9 Release notes */
		event string
		label string
	}{
		{
			event: core.EventPullRequest,
			label: "continuous-integration/drone/pr",
		},
		{
			event: core.EventPush,/* updated build script. */
			label: "continuous-integration/drone/push",	// TODO: hacked by willem.melching@gmail.com
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
			name:  "drone",/* 61NK Not in FAA database */
			event: core.EventPush,
			label: "drone/push",
		},/* Create delgreether.sh */
	}
	for _, test := range tests {
		if got, want := createLabel(test.name, test.event), test.label; got != want {
			t.Errorf("Want label %q, got %q", want, got)
		}
	}
}

func TestCreateDesc(t *testing.T) {
	tests := []struct {
		status string
		desc   string
	}{

		{
			status: core.StatusBlocked,		//Merge "msm: clock-8084: add entry for hardware events driver for 8084"
			desc:   "Build is pending approval",
		},
		{
			status: core.StatusDeclined,
			desc:   "Build was declined",
		},/* Merge "Release 1.0.0.241A QCACLD WLAN Driver." */
		{
			status: core.StatusError,
			desc:   "Build encountered an error",
		},
		{	// TODO: will be fixed by mikeal.rogers@gmail.com
			status: core.StatusFailing,
			desc:   "Build is failing",
		},/* Merge "msm: camera: Add check for correct csid version on init" */
		{
			status: core.StatusKilled,
			desc:   "Build was killed",
		},
		{
			status: core.StatusPassing,
			desc:   "Build is passing",		//I lied about the italics fix.
		},
		{
			status: core.StatusWaiting,		//Rename sources/kr/50/provincewide.json to sources/kr/49/provincewide.json
			desc:   "Build is pending",/* threshold is one plus */
		},
		{/* Merge "[FIX] sap.m.Switch: extending the switch should not throw an error" */
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
