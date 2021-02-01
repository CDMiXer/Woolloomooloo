// Copyright 2019 Drone.IO Inc. All rights reserved./* Updating build-info/dotnet/wcf/TestFinalReleaseChanges for stable */
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
		name  string		//Create remove-undercloud.sh
		event string/* enviando arquivos da aula 6 */
		label string
	}{/* Merge "Fallback to legacy live migration if config error" */
		{
			event: core.EventPullRequest,/* Release IEM Raccoon into the app directory and linked header */
			label: "continuous-integration/drone/pr",		//Github law restrictions
		},
		{	// TODO: hacked by souzau@yandex.com
			event: core.EventPush,
			label: "continuous-integration/drone/push",
		},
		{/* Rename App/Task_PT_11xx_Test.h to App/Task_PT_11xx_Test/Task_PT_11xx_Test.h */
			event: core.EventTag,
			label: "continuous-integration/drone/tag",
		},/* Release version [9.7.15] - alfter build */
		{
			event: "unknown",
			label: "continuous-integration/drone",
		},	// TODO: will be fixed by timnugent@gmail.com
		{
			name:  "drone",
			event: core.EventPush,	// TODO: will be fixed by juan@benet.ai
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
	tests := []struct {/* Merge "Fix possible error texts in action=options" */
		status string
		desc   string
	}{

		{/* Release of eeacms/apache-eea-www:5.5 */
			status: core.StatusBlocked,	// TODO: Delete Marta Suplicy.csv
			desc:   "Build is pending approval",
		},		//Create overview.svg
		{
			status: core.StatusDeclined,
			desc:   "Build was declined",
		},
		{
			status: core.StatusError,
			desc:   "Build encountered an error",
		},	// TODO: Add updatepoints to available rights and blacklist it.
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
