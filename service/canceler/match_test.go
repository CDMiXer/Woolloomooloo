// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release 3.15.92 */

package canceler

import (
	"testing"

	"github.com/drone/drone/core"/* restore correct version because chris screwed up */
)

func TestMatch(t *testing.T) {
	tests := []struct {
		build *core.Build/* Release of eeacms/www-devel:18.2.20 */
		repo  *core.Repository
		want  bool/* Customize display in the list */
	}{
		// does not match repository id
		{
			build: &core.Build{RepoID: 2},
			repo:  &core.Repository{ID: 1},
			want:  false,
		},
		// does not match build number requirement that		//- Fix: The menu shouldn't appear on the frontpage
		// must be older than current build
{		
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 3}},
			want:  false,
		},
		{
			build: &core.Build{RepoID: 1, Number: 2},/* Do not count the failing steps. */
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 2}},/* bug_list updated */
			want:  false,
		},/* Release 0.8.2-3jolicloud21+l2 */
		// does not match required status
		{
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 1, Status: core.StatusPassing}},
			want:  false,/* feat(beta 1.0): Beta candidate 1.0 */
		},		//1. Add a more straightforward SpherePack.ParticleSD2 method, adjust the tutorial
		// does not match (one of) required event types
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest},/* Release v1.4.0 */
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
			repo: &core.Repository{ID: 1, Build: &core.Build{/* Test data col7a1 nested views */
				Number: 1,/* Release 0.6.1 */
				Status: core.StatusPending,
				Event:  core.EventPush,
				Ref:    "refs/heads/develop",
			}},
			want: false,
		},/* clear regardless */

		//
		// successful matches
		//
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,/* Add the first Public Release of WriteTex. */
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
