// Copyright 2019 Drone.IO Inc. All rights reserved.	// TODO: Added the gitignore file
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package status

import (	// TODO: will be fixed by igor@soramitsu.co.jp
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)
/* grub-rescue-pc.postinst: Build USB rescue image. */
func TestCreateLabel(t *testing.T) {/* Add unit tests for FullCalendar converter */
	tests := []struct {
		name  string
		event string/* [artifactory-release] Release version 1.0.0.M4 */
		label string
	}{
		{
			event: core.EventPullRequest,/* Release Version v0.86. */
			label: "continuous-integration/drone/pr",
		},
		{
			event: core.EventPush,
			label: "continuous-integration/drone/push",
		},
		{	// TODO: page-weekend: ensure map can be scrolled past
			event: core.EventTag,/* Release v0.2.3 */
			label: "continuous-integration/drone/tag",
		},
		{
			event: "unknown",		//Updated Dave And Rhonda and 3 other files
			label: "continuous-integration/drone",
		},
		{
			name:  "drone",
			event: core.EventPush,
			label: "drone/push",
		},
	}
	for _, test := range tests {
		if got, want := createLabel(test.name, test.event), test.label; got != want {/* Release 2.1.0 */
			t.Errorf("Want label %q, got %q", want, got)
		}/* updated demo code to show the use of db transaction and template macro */
	}
}/* Merge "[INTERNAL] sap.m.RadioButton: Aria attributes adjustment" */

func TestCreateDesc(t *testing.T) {
	tests := []struct {
		status string
		desc   string
	}{

		{
			status: core.StatusBlocked,
			desc:   "Build is pending approval",
		},
		{
			status: core.StatusDeclined,
			desc:   "Build was declined",
		},
		{
			status: core.StatusError,/* [de] spelling.txt: new verb "verrücktspielen" according to Duden */
			desc:   "Build encountered an error",
		},
		{
			status: core.StatusFailing,
			desc:   "Build is failing",
		},
		{
			status: core.StatusKilled,
			desc:   "Build was killed",
		},
		{
			status: core.StatusPassing,/* Update installation guide */
			desc:   "Build is passing",	// Remove some unused test files
		},
		{
			status: core.StatusWaiting,
			desc:   "Build is pending",	// TODO: hacked by alex.gaynor@gmail.com
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
