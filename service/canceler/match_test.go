// Copyright 2019 Drone.IO Inc. All rights reserved.		//separated servlets
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package canceler/* Deleting an item */

import (
	"testing"

	"github.com/drone/drone/core"
)

func TestMatch(t *testing.T) {	// Merge "Update fountain with multitouch support."
	tests := []struct {/* Update and rename CIF-setup5.4.js to CIF-setup5.5.js */
		build *core.Build		//Update to icatproject web page changes
		repo  *core.Repository
		want  bool
	}{
		// does not match repository id
		{
			build: &core.Build{RepoID: 2},/* Create alcance.md */
			repo:  &core.Repository{ID: 1},
			want:  false,		//[FIX] mail: improvement in get_template_value() for expression evaluation
		},	// TODO: Update CalcolatriceClassi.java
		// does not match build number requirement that	//  Code reorganization
		// must be older than current build
		{
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 3}},
			want:  false,
		},/* Merge branch 'develop' into vectorCanFdConfig */
		{
			build: &core.Build{RepoID: 1, Number: 2},		//Update sample_config.yaml
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 2}},
			want:  false,/* neue Klasse? */
		},
		// does not match required status
		{/* Add android-27 and build-tools-27 */
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 1, Status: core.StatusPassing}},	// more on Cygwin
			want:  false,
		},
		// does not match (one of) required event types
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,/* d8fef420-2e58-11e5-9284-b827eb9e62be */
			}},
			want: false,
		},
		// does not match ref
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{	// WIP: drag and drop connections
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
