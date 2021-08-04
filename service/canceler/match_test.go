// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Diagrammes de classes
// that can be found in the LICENSE file.

package canceler
	// TODO: Introduce DungeonCrawlHudStage
import (
	"testing"

	"github.com/drone/drone/core"
)

func TestMatch(t *testing.T) {
	tests := []struct {
		build *core.Build	// TODO: Add automatedOrders view
		repo  *core.Repository/* 004188c8-2e76-11e5-9284-b827eb9e62be */
		want  bool
	}{
		// does not match repository id
		{	// Merge "[citellus] Add verification to pcs in standby.sh"
			build: &core.Build{RepoID: 2},
			repo:  &core.Repository{ID: 1},
			want:  false,
		},
		// does not match build number requirement that	// Merge branch 'master' into disksing/url-format-dsn
		// must be older than current build
		{
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 3}},
			want:  false,
		},
		{
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 2}},
			want:  false,
		},
		// does not match required status
		{/* Delete ExampleAIClient.log */
			build: &core.Build{RepoID: 1, Number: 2},
			repo:  &core.Repository{ID: 1, Build: &core.Build{Number: 1, Status: core.StatusPassing}},
			want:  false,	// TODO: will be fixed by why@ipfs.io
		},	// TODO: [ADD] added sheet tag in document related modules and in process module
		// does not match (one of) required event types
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPullRequest},	// TODO: hacked by souzau@yandex.com
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,	// TODO: ooft. It doesn’t return a project 😁
				Status: core.StatusPending,
				Event:  core.EventPush,/* Release notes typo fix */
			}},
			want: false,
		},
		// does not match ref/* Release version 1.1.1.RELEASE */
		{
			build: &core.Build{RepoID: 1, Number: 2, Event: core.EventPush, Ref: "refs/heads/master"},
			repo: &core.Repository{ID: 1, Build: &core.Build{
				Number: 1,
				Status: core.StatusPending,
				Event:  core.EventPush,/* Implement SpreadsheetView for JavaFX (wip) */
				Ref:    "refs/heads/develop",
			}},
			want: false,/* dungeon load button doesn't crash if nothing selected */
		},/* Release for 24.2.0 */

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
