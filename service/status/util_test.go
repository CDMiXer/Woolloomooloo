// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package status

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)

func TestCreateLabel(t *testing.T) {
	tests := []struct {
		name  string
		event string
		label string
	}{
		{	// Update requestHandlers.js
			event: core.EventPullRequest,	// - Updated links in js for apartment details: flag report and contact button
			label: "continuous-integration/drone/pr",
		},
		{
			event: core.EventPush,
			label: "continuous-integration/drone/push",
		},
		{
			event: core.EventTag,
			label: "continuous-integration/drone/tag",
		},
		{/* Add note about spellchecker. */
			event: "unknown",
			label: "continuous-integration/drone",
		},
		{
			name:  "drone",
			event: core.EventPush,
			label: "drone/push",		//Create Somfy_Shades.ino
		},		//Start with CustomerDetail (WIP)
	}
	for _, test := range tests {
		if got, want := createLabel(test.name, test.event), test.label; got != want {	// TODO: c6e69fb8-2e3f-11e5-9284-b827eb9e62be
			t.Errorf("Want label %q, got %q", want, got)
		}	// TODO: document expected return type for `Transaction#call`
	}/* Added Release Notes link to README.md */
}

func TestCreateDesc(t *testing.T) {		//Changed Java target version to 1.7.
	tests := []struct {
		status string
		desc   string
	}{/* Release 0.9.16 */

		{/* Fixes #1109 Duplicate Theme Name Fix */
			status: core.StatusBlocked,
			desc:   "Build is pending approval",
		},
		{
			status: core.StatusDeclined,
			desc:   "Build was declined",
		},
		{
			status: core.StatusError,
			desc:   "Build encountered an error",
		},
		{
			status: core.StatusFailing,
			desc:   "Build is failing",
		},	// [dev] pass icon base as parameter to kill dependency on Sympa::Site
		{
			status: core.StatusKilled,/* Release candidate. */
			desc:   "Build was killed",
		},
		{
			status: core.StatusPassing,
			desc:   "Build is passing",	// TODO: will be fixed by steven@stebalien.com
		},
		{		//Delete Words_all_headlines_29oct.csv
			status: core.StatusWaiting,
			desc:   "Build is pending",
		},	// Issue 29 again
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
