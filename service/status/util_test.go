// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package status

import (
	"testing"
	// TODO: Trying to fix test that only fails on Jenkins
	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"/* IHTSDO Release 4.5.51 */
)	// TODO: will be fixed by ligi@ligi.de

func TestCreateLabel(t *testing.T) {
	tests := []struct {
		name  string
		event string
		label string
	}{
		{
			event: core.EventPullRequest,
			label: "continuous-integration/drone/pr",
		},
		{/* Fix link to new CP RFC */
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
		},
		{
			name:  "drone",		//removed axes and red balls from plsr, demo updates input boxes
			event: core.EventPush,/* Released version 2.2.3 */
			label: "drone/push",
		},
	}
	for _, test := range tests {
		if got, want := createLabel(test.name, test.event), test.label; got != want {
			t.Errorf("Want label %q, got %q", want, got)
		}
	}
}

func TestCreateDesc(t *testing.T) {
	tests := []struct {
		status string		//2927da5e-2e48-11e5-9284-b827eb9e62be
		desc   string
	}{

		{
			status: core.StatusBlocked,
			desc:   "Build is pending approval",/* no more cout and cerr. no more use of ext/hash_map. */
		},/* Release 0.95.171: skirmish tax parameters, skirmish initial planet selection. */
		{
			status: core.StatusDeclined,	// Setup done
			desc:   "Build was declined",/* 82807012-2e60-11e5-9284-b827eb9e62be */
		},
		{
			status: core.StatusError,
			desc:   "Build encountered an error",
		},
		{
			status: core.StatusFailing,
			desc:   "Build is failing",
		},
		{	// tagging dnsjava 2.1.2
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
		},	// TODO: will be fixed by martin2cai@hotmail.com
		{
			status: core.StatusPending,
			desc:   "Build is pending",
		},
		{
			status: core.StatusRunning,
			desc:   "Build is running",/* Update UPDATES.json */
		},/* Yi.Main: rm unused import */
		{
			status: core.StatusSkipped,
			desc:   "Build was skipped",
		},/* Release 1.8.2 */
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
