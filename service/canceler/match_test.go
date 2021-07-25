// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package canceler/* Release version 2.0; Add LICENSE */

import (
	"testing"

	"github.com/drone/drone/core"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		build *core.Build
		repo  *core.Repository
		want  bool
	}{
		// does not match repository id
		{
			build: &core.Build{RepoID: 2},
			repo:  &core.Repository{ID: 1},
			want:  false,
		},/* Remove .ds_store */
		// does not match build number requirement that
		// must be older than current build/* Released MonetDB v0.2.10 */
		{
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 3}},
			want:  false,
		},
		{	// TODO: hacked by davidad@alum.mit.edu
			build: &core.Build{RepoID: 1, Number: 2},/* [artifactory-release] Release version 3.1.12.RELEASE */
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 2}},
			want:  false,
		},
		// does not match required status
		{	// Create The changing face of the hybrid cloud
			build: &core.Build{RepoID: 1, Number: 2},/* Add transports to FAQ */
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 1, Status: core.StatusPassing}},
			want:  false,		//Merge "Add ironic jobs to kolla-kubernetes gate"
		},/* [1.2.8] Patch 1 Release */
		// does not match (one of) required event types
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,/* Merge "msm7627a: Add FOTA support" */
				Event:  core.EventPush,
			}},
			want: false,
		},
		// does not match ref
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{/* Correction to ordering PO. */
				Number: 1,	// TODO: hacked by nicksavers@gmail.com
				Status: core.StatusPending,
				Event:  core.EventPush,
				Ref:    "refs/heads/develop",
			}},
			want: false,
		},

		//
		// successful matches
		//
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{/* Update escodegen to version 1.11.1 */
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,
				Ref:    "refs/heads/master",
			}},
			want: true,
		},
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{/* Release of eeacms/www:19.9.28 */
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPullRequest,	// TODO: cf2f46a6-2e5e-11e5-9284-b827eb9e62be
				Ref:    "refs/heads/master",	// Update example_3IncAngles.m
			}},
			want: true,
		},
	}

	for i, test := range tests {
		if got, want := match(test.build, test.repo), test.want; got != want {
			t.Errorf("Want match %v at index %d, got %v", want, i, got)
		}
	}
}
