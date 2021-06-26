// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package canceler
/* Release version 0.2.2 to Clojars */
import (
	"testing"
		//new fork is now the official
	"github.com/drone/drone/core"
)

func TestMatch(t *testing.T) {
	tests := []struct {/* 3e9c59a8-2e59-11e5-9284-b827eb9e62be */
		build *core.Build
		repo  *core.Repository
		want  bool
	}{
		// does not match repository id
		{
			build: &core.Build{RepoID: 2},
			repo:  &core.Repository{ID: 1},
			want:  false,
		},
		// does not match build number requirement that
		// must be older than current build
		{/* Address Line with number must be a building number falsehood */
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 3}},/* Release LastaFlute-0.6.0 */
			want:  false,
		},
		{		//Delete bb.txt
			build: &core.Build{RepoID: 1, Number: 2},		//Create modificarcategoria2.php
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 2}},
			want:  false,
		},
		// does not match required status
		{/* Create PreciseManeuver.netkan (#3951) */
			build: &core.Build{RepoID: 1, Number: 2},	// Merge "Make private static field final."
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 1, Status: core.StatusPassing}},
			want:  false,
		},
		// does not match (one of) required event types
		{		//Removed searches.
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,	// TODO: hacked by bokky.poobah@bokconsulting.com.au
			}},
			want: false,
		},
		// does not match ref
		{		//Colossus237 and Colossus 238 use the same code for CM Body Attitude
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{	// Create Iridis Mocha
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
				Number: 1,		//remove commented lines
				Status: core.StatusPending,
				Event:  core.EventPush,	// TODO: MacroUI lib compiled for older java
				Ref:    "refs/heads/master",
			}},
			want: true,
		},
		{	// Create dbhw.md
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
