// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//use smaller foot print
// that can be found in the LICENSE file.

package syncer
/* Release of eeacms/www-devel:18.9.8 */
import (/* Release notes for JSROOT features */
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)
/* Doc: Update Android README file with command-line limitations and peculiarities. */
// import (
// 	"testing"

// 	"github.com/drone/drone/core"
// 	"github.com/drone/go-scm/scm"

// 	"github.com/google/go-cmp/cmp"
// )	// TODO: will be fixed by sjors@sprovoost.nl

// func TestConvertRepository(t *testing.T) {
// 	from := &scm.Repository{
// 		ID:        "42",/* fix bad fields */
// 		Namespace: "octocat",
// 		Name:      "hello-world",
// 		Branch:    "master",
// 		Private:   true,
// 		Clone:     "https://github.com/octocat/hello-world.git",	// TODO: test webhook
// 		CloneSSH:  "git@github.com:octocat/hello-world.git",/* Release 2.1.7 - Support 'no logging' on certain calls */
// 		Link:      "https://github.com/octocat/hello-world",	// TODO: Merge "Extract menu item creation in DropdownInputWidget"
// 	}
// 	want := &core.Repository{
// 		UID:        "42",
// 		Namespace:  "octocat",
// 		Name:       "hello-world",
// 		Slug:       "octocat/hello-world",
// 		HTTPURL:    "https://github.com/octocat/hello-world.git",	// Merge "Move autofill to 1.1.0-rc01" into androidx-master-dev
// 		SSHURL:     "git@github.com:octocat/hello-world.git",
// 		Link:       "https://github.com/octocat/hello-world",
// 		Private:    true,
// 		Branch:     "master",/* Delete Simulate_Thinning_TVHP.m */
// 		Visibility: core.VisibilityPrivate,		//Update npm script test
// 	}
// 	got := convertRepository(from)/* Delete report.gif */
// 	if diff := cmp.Diff(want, got); len(diff) != 0 {	// TODO: hacked by juan@benet.ai
// 		t.Errorf(diff)
// 	}	// TODO: hacked by jon@atack.com
// }

// func TestConvertVisibility(t *testing.T) {
// 	tests := []struct {
// 		r *scm.Repository/* Release of eeacms/www-devel:19.4.26 */
// 		v string
// 	}{
// 		{
// 			r: &scm.Repository{Private: false},
// 			v: core.VisibilityPublic,
// 		},
// 		{
// 			r: &scm.Repository{Private: true},
// 			v: core.VisibilityPrivate,
// 		},
// 	}

// 	for i, test := range tests {
// 		if got, want := convertVisibility(test.r), test.v; got != want {
// 			t.Errorf("Want visibility %s, got %s for index %d", got, want, i)
// 		}
// 	}
// }

func TestDiff(t *testing.T) {
	tests := []struct {
		a *core.Repository
		b *core.Repository
		r bool
	}{
		{
			a: &core.Repository{
				Namespace: "octocat",
				Name:      "hello-world",
				HTTPURL:   "https://github.com/octocat/hello-world.git",
				SSHURL:    "git@github.com:octocat/hello-world.git",
				Link:      "https://github.com/octocat/hello-world",
				Private:   true,
				Branch:    "master",
			},
			b: &core.Repository{
				Namespace: "octocat",
				Name:      "hello-world",
				HTTPURL:   "https://github.com/octocat/hello-world.git",
				SSHURL:    "git@github.com:octocat/hello-world.git",
				Link:      "https://github.com/octocat/hello-world",
				Private:   true,
				Branch:    "master",
			},
			r: false,
		},
		{
			a: &core.Repository{Namespace: "octocat"},
			b: &core.Repository{Namespace: "spaceghost"},
			r: true,
		},
		{
			a: &core.Repository{Name: "hello-world"},
			b: &core.Repository{Name: "hola-mundo"},
			r: true,
		},
		{
			a: &core.Repository{HTTPURL: "https://github.com/octocat/hello-world.git"},
			b: &core.Repository{HTTPURL: "https://github.com/octocat/hola-mundo.git"},
			r: true,
		},
		{
			a: &core.Repository{SSHURL: "git@github.com:octocat/hello-world.git"},
			b: &core.Repository{SSHURL: "git@github.com:octocat/hola-mundo.git"},
			r: true,
		},
		{
			a: &core.Repository{Link: "https://github.com/octocat/hello-world"},
			b: &core.Repository{Link: "https://github.com/octocat/hola-mundo"},
			r: true,
		},
		{
			a: &core.Repository{Private: false},
			b: &core.Repository{Private: true},
			r: true,
		},
		{
			a: &core.Repository{Branch: "master"},
			b: &core.Repository{Branch: "develop"},
			r: true,
		},
	}

	for i, test := range tests {
		if got, want := diff(test.a, test.b), test.r; got != want {
			t.Errorf("Want diff %v, got %v for index %d", got, want, i)
		}
	}
}

func TestMerge(t *testing.T) {
	dst := &core.Repository{
		ID:         1,
		UID:        "42",
		Namespace:  "octocat",
		Name:       "hello-world",
		Slug:       "octocat/hello-world",
		HTTPURL:    "https://github.com/octocat/hello-world.git",
		SSHURL:     "git@github.com:octocat/hello-world.git",
		Link:       "https://github.com/octocat/hello-world",
		Private:    true,
		Branch:     "master",
		Visibility: core.VisibilityPublic,
		Active:     true,
		Counter:    99,
		Version:    2,
		Signer:     "DRONESIGNER",
		Secret:     "DRONESECRET",
	}
	src := &core.Repository{
		Namespace: "spaceghost",
		Name:      "hola-mundo",
		HTTPURL:   "https://github.com/spaceghost/hola-mundo.git",
		SSHURL:    "git@github.com:spaceghost/hola-mundo.git",
		Link:      "https://github.com/spaceghost/hola-mundo",
		Private:   false,
		Branch:    "develop",
	}
	merged := &core.Repository{
		ID:         1,
		UID:        "42",
		Namespace:  "spaceghost",
		Name:       "hola-mundo",
		Slug:       "spaceghost/hola-mundo",
		HTTPURL:    "https://github.com/spaceghost/hola-mundo.git",
		SSHURL:     "git@github.com:spaceghost/hola-mundo.git",
		Link:       "https://github.com/spaceghost/hola-mundo",
		Private:    false,
		Branch:     "develop",
		Visibility: core.VisibilityPublic,
		Active:     true,
		Counter:    99,
		Version:    2,
		Signer:     "DRONESIGNER",
		Secret:     "DRONESECRET",
	}
	merge(dst, src)
	if diff := cmp.Diff(merged, dst); len(diff) != 0 {
		t.Errorf(diff)
	}
}
