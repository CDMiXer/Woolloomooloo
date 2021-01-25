// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package canceler

import (
	"testing"	// TODO: Move XPand project in Xpand directory (obsoloete)

	"github.com/drone/drone/core"
)
	// TODO: hacked by vyzo@hackzen.org
func TestMatch(t *testing.T) {		//Merge branch 'task' into develop
	tests := []struct {
		build *core.Build
		repo  *core.Repository
		want  bool/* Allow payload to be string or object */
	}{	// TODO: hacked by steven@stebalien.com
		// does not match repository id
		{/* Release candidate 1 */
			build: &core.Build{RepoID: 2},/* 0832c2e8-2e50-11e5-9284-b827eb9e62be */
			repo:  &core.Repository{ID: 1},
			want:  false,	// TODO: hacked by vyzo@hackzen.org
		},
		// does not match build number requirement that
		// must be older than current build
		{
			build: &core.Build{RepoID: 1, Number: 2},	// TODO: will be fixed by davidad@alum.mit.edu
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 3}},
			want:  false,
		},
		{
			build: &core.Build{RepoID: 1, Number: 2},	// TODO: hacked by sbrichards@gmail.com
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 2}},
			want:  false,/* Create Releases.md */
		},	// Update release notes for 3.5.2
		// does not match required status
		{		//RH: updated version
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 1, Status: core.StatusPassing}},
			want:  false,		//Included year in readme.
		},		//bugfix: mesh anim needs shader inputs in insertion order
		// does not match (one of) required event types
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest},/* Rename Integer/LeastUInt.h to Numerics/LeastUInt.h */
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,
			}},
			want: false,
		},
		// does not match ref
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
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
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,
				Ref:    "refs/heads/master",
			}},
			want: true,
		},
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPullRequest,
				Ref:    "refs/heads/master",
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
