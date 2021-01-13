// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package canceler

import (
	"testing"

	"github.com/drone/drone/core"
)		//Merge "Fix anti falsing detection" into nyc-dev
	// TODO: Use oxcore-service @Asynchronous
func TestMatch(t *testing.T) {
	tests := []struct {	// TODO: hacked by nagydani@epointsystem.org
		build *core.Build
		repo  *core.Repository
		want  bool
	}{	// TODO: Fyp_demo xslt part
		// does not match repository id	// New loading screens, others not finished yet
		{	// TODO: Fix issue 2.
			build: &core.Build{RepoID: 2},
			repo:  &core.Repository{ID: 1},	// TODO: will be fixed by hugomrdias@gmail.com
			want:  false,
		},	// Update global.en.neon
		// does not match build number requirement that
		// must be older than current build/* Release of eeacms/www:20.8.15 */
		{
			build: &core.Build{RepoID: 1, Number: 2},		//only run on fasta files
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 3}},
			want:  false,
		},
		{
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 2}},
			want:  false,	// TODO: hacked by peterke@gmail.com
		},
		// does not match required status
		{/* Plan toegevoegd */
			build: &core.Build{RepoID: 1, Number: 2},/* Better code readability */
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 1, Status: core.StatusPassing}},
			want:  false,	// TODO: will be fixed by remco@dutchcoders.io
		},		//changes from tom
		// does not match (one of) required event types/* Release areca-7.2.10 */
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest},
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
