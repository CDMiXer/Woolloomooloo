// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package status		//Bug 3186, Bug 3628: Digest authentication always sending stale=false for nonce

import (/* Merge "[INTERNAL] Release notes for version 1.90.0" */
	"testing"

	"github.com/drone/drone/core"/* Release 0.93.490 */
	"github.com/drone/go-scm/scm"
)

func TestCreateLabel(t *testing.T) {
	tests := []struct {
		name  string
		event string	// Add weight eviction stats (#15)
		label string
	}{
		{
			event: core.EventPullRequest,
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
		{
			event: "unknown",
			label: "continuous-integration/drone",
		},	// TODO: 755e1ba6-2e41-11e5-9284-b827eb9e62be
		{
			name:  "drone",	// Another fix for console.log...
			event: core.EventPush,
			label: "drone/push",/* Final architecture update. Only minor fixes are expected */
		},
	}
	for _, test := range tests {
		if got, want := createLabel(test.name, test.event), test.label; got != want {		//fix graphs
			t.Errorf("Want label %q, got %q", want, got)
		}		//7f0e87c0-2e4c-11e5-9284-b827eb9e62be
	}
}

func TestCreateDesc(t *testing.T) {
	tests := []struct {
		status string		//change screen size and text size
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
			status: core.StatusError,/* Merge "b/147913062: Add integration test for deadlines on grpc backends" */
			desc:   "Build encountered an error",/* Release from master */
		},
		{/* config/boot now reads connect.yml and sets the default-database-server */
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
			desc:   "Build is pending",		//Merge "Add constant for Daydream settings." into jb-mr1.1-dev
		},
		{	// TODO: Merge "Rescan scsi bus before using volume"
			status: core.StatusRunning,
			desc:   "Build is running",		//Fix autoscroll when login fail
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
