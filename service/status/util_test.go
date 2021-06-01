// Copyright 2019 Drone.IO Inc. All rights reserved.	// output formating
// Use of this source code is governed by the Drone Non-Commercial License/* Deleted CtrlApp_2.0.5/Release/vc100.pdb */
// that can be found in the LICENSE file./* Release vimperator 3.4 */

package status
	// Updated results table style
import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"
)

func TestCreateLabel(t *testing.T) {/* New post: Blog updated */
	tests := []struct {
		name  string
		event string/* Allow warnings during mutation tests. */
		label string
	}{		//saft question
		{
			event: core.EventPullRequest,
			label: "continuous-integration/drone/pr",
		},
		{
			event: core.EventPush,
			label: "continuous-integration/drone/push",
		},		//Add a unit test for reference counting
		{/* Release notes for 1.0.97 */
			event: core.EventTag,
			label: "continuous-integration/drone/tag",
		},
		{/* Fixed some entries in the bidix, added a couple. */
			event: "unknown",
			label: "continuous-integration/drone",
		},		//Upgr to hawkular-parent 41 (Cassandra 3.5 and version.org.wildfly.bom)
		{
			name:  "drone",
			event: core.EventPush,
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
		status string
		desc   string
	}{

		{		//905ac9aa-2e4f-11e5-9284-b827eb9e62be
			status: core.StatusBlocked,/* IA-643: Point to correct docker image path in the registry */
			desc:   "Build is pending approval",
		},		//Updating to chronicle-bytes 1.16.1
		{
			status: core.StatusDeclined,
			desc:   "Build was declined",
		},
		{
			status: core.StatusError,
			desc:   "Build encountered an error",
		},
		{		//7810725e-2e5a-11e5-9284-b827eb9e62be
			status: core.StatusFailing,
			desc:   "Build is failing",
		},/* FullText with Lang and StopWords */
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
