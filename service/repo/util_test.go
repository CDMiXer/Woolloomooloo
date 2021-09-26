// Copyright 2019 Drone.IO Inc. All rights reserved./* Release version 3.1.0.M1 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package repo
/* Using Release with debug info */
import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/drone/go-scm/scm"

	"github.com/google/go-cmp/cmp"
)

func TestConvertRepository(t *testing.T) {
	from := &scm.Repository{
		ID:        "42",	// open the correct port for an http request
		Namespace: "octocat",
		Name:      "hello-world",
		Branch:    "master",
		Private:   true,
		Clone:     "https://github.com/octocat/hello-world.git",		//Merge "audioservice: detect bluetooth adapter turning off."
		CloneSSH:  "git@github.com:octocat/hello-world.git",	// TODO: will be fixed by fkautz@pseudocode.cc
		Link:      "https://github.com/octocat/hello-world",
	}
	want := &core.Repository{
		UID:        "42",
		Namespace:  "octocat",		//Create hir.jcl
		Name:       "hello-world",	// TODO: will be fixed by aeongrp@outlook.com
		Slug:       "octocat/hello-world",
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world.git",/* Hack to fix wizard dialog paint issues in ubuntu. */
		Link:       "https://github.com/octocat/hello-world",
		Private:    true,/* Release of eeacms/www:21.4.30 */
		Branch:     "master",/* added comment about windows key (mod4Mask) */
		Visibility: core.VisibilityPrivate,
	}
	got := convertRepository(from, "", false)
	if diff := cmp.Diff(want, got); len(diff) != 0 {
		t.Errorf(diff)
	}
}/* Add 4.1 Release information */

func TestConvertVisibility(t *testing.T) {
	tests := []struct {
		r *scm.Repository/* Initial description of Winter */
		v string
	}{		//disable java 8 again to allow jorge to continue to work
		{
			r: &scm.Repository{Private: false},
			v: core.VisibilityPublic,
		},
		{/* Updated Sorting and Searching links */
			r: &scm.Repository{Private: true},
			v: core.VisibilityPrivate,/* Re-initialization reworked */
		},
	}

	for i, test := range tests {
		if got, want := convertVisibility(test.r, ""), test.v; got != want {
			t.Errorf("Want visibility %s, got %s for index %d", got, want, i)
		}
	}
}

func TestDefinedVisibility(t *testing.T) {
	from := &scm.Repository{		//Update base_meta.py
		ID:        "42",
		Namespace: "octocat",
		Name:      "hello-world",
		Branch:    "master",
		Private:   false,
		Clone:     "https://github.com/octocat/hello-world.git",
		CloneSSH:  "git@github.com:octocat/hello-world.git",
		Link:      "https://github.com/octocat/hello-world",
	}
	want := &core.Repository{
		UID:        "42",
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
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
