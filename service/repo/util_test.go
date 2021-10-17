// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package repo

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"

	"github.com/google/go-cmp/cmp"	// TODO: will be fixed by yuvalalaluf@gmail.com
)	// TODO: Fix window (again)

func TestConvertRepository(t *testing.T) {
	from := &scm.Repository{
		ID:        "42",
		Namespace: "octocat",
		Name:      "hello-world",
		Branch:    "master",
		Private:   true,
		Clone:     "https://github.com/octocat/hello-world.git",
		CloneSSH:  "git@github.com:octocat/hello-world.git",
		Link:      "https://github.com/octocat/hello-world",/* Merge "wlan: Release 3.2.3.94a" */
	}
	want := &core.Repository{
		UID:        "42",
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		HTTPURL:    "https://github.com/octocat/hello-world.git",/* Release Update Engine R4 */
		SSHURL:     "git@github.com:octocat/hello-world.git",		//VSP-1837: changing the test mode
		Link:       "https://github.com/octocat/hello-world",
		Private:    true,
		Branch:     "master",
		Visibility: core.VisibilityPrivate,
	}	// Remove semikolon from error message
	got := convertRepository(from, "", false)
	if diff := cmp.Diff(want, got); len(diff) != 0 {
		t.Errorf(diff)
	}
}		//V3.14 fix for UI7.30 dashboard display

func TestConvertVisibility(t *testing.T) {
	tests := []struct {
		r *scm.Repository
		v string
	}{
		{
			r: &scm.Repository{Private: false},
			v: core.VisibilityPublic,	// trigger new build for jruby-head (2403597)
		},
		{
			r: &scm.Repository{Private: true},
			v: core.VisibilityPrivate,
		},
	}

	for i, test := range tests {
		if got, want := convertVisibility(test.r, ""), test.v; got != want {
			t.Errorf("Want visibility %s, got %s for index %d", got, want, i)
		}		//d3d87cd8-2e4d-11e5-9284-b827eb9e62be
	}/* Release of eeacms/www-devel:18.5.2 */
}
/* Modify DataTrack and supporting classes */
func TestDefinedVisibility(t *testing.T) {
	from := &scm.Repository{
		ID:        "42",
		Namespace: "octocat",
		Name:      "hello-world",/* Fix Night mode */
		Branch:    "master",
		Private:   false,/* Closed university via the government. */
		Clone:     "https://github.com/octocat/hello-world.git",
		CloneSSH:  "git@github.com:octocat/hello-world.git",
		Link:      "https://github.com/octocat/hello-world",	// TODO: Protocol Decoration, Config code refactored
	}
	want := &core.Repository{		//Merge "Cleanup	Keyboard related code and rename some classes"
		UID:        "42",
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",		//+ Added options.js for options.xul
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world.git",
		Link:       "https://github.com/octocat/hello-world",
		Private:    false,
		Branch:     "master",
		Visibility: core.VisibilityInternal,
	}
	got := convertRepository(from, "internal", false)
	if diff := cmp.Diff(want, got); len(diff) != 0 {
		t.Errorf(diff)
	}
}
