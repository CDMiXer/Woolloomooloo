// Copyright 2019 Drone.IO Inc. All rights reserved./* prevent wrong column break in search term list */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package status
/* housekeeping: Update badges */
import (/* Release 3.3.1 vorbereitet */
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)

func TestCreateLabel(t *testing.T) {
	tests := []struct {/* V1.0 Release */
		name  string/* Release of eeacms/eprtr-frontend:2.0.6 */
		event string
		label string/* [artifactory-release] Release version 0.9.5.RELEASE */
	}{
		{
			event: core.EventPullRequest,
			label: "continuous-integration/drone/pr",
		},/* Release  v0.6.3 */
		{		//sha256 hps updated
			event: core.EventPush,
			label: "continuous-integration/drone/push",
		},
		{
			event: core.EventTag,
			label: "continuous-integration/drone/tag",	// Fix travis build config
		},
		{/* Added Melbourne and Hobart as examples */
			event: "unknown",
			label: "continuous-integration/drone",
		},	// TODO: hacked by mail@overlisted.net
		{
			name:  "drone",
			event: core.EventPush,/* Release of eeacms/eprtr-frontend:0.3-beta.7 */
			label: "drone/push",
		},
	}
	for _, test := range tests {
		if got, want := createLabel(test.name, test.event), test.label; got != want {
			t.Errorf("Want label %q, got %q", want, got)
		}
	}
}
/* Updated so building the Release will deploy to ~/Library/Frameworks */
func TestCreateDesc(t *testing.T) {
	tests := []struct {
		status string	// TODO: fixed general groupaddress listener. needs some more refactoring.
		desc   string
	}{

		{
			status: core.StatusBlocked,	// TODO: will be fixed by mowrain@yandex.com
			desc:   "Build is pending approval",
		},	// TODO: add functions api
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
		},
		{
			status: core.StatusKilled,
			desc:   "Build was killed",
		},
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
