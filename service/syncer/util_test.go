// Copyright 2019 Drone.IO Inc. All rights reserved.	// Re-add User#role, which now returns the first user role.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Release v3.2.2 compatiable with joomla 3.2.2 */
package syncer

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)/* Release version 2.0.0.RELEASE */
	// [Issue #91] fix throw and allow JSScreenWrapper
// import (
// 	"testing"

// 	"github.com/drone/drone/core"
// 	"github.com/drone/go-scm/scm"/* * Release 2.2.5.4 */

// 	"github.com/google/go-cmp/cmp"
// )
	// TODO: Ajout macro G. glabrum
// func TestConvertRepository(t *testing.T) {
// 	from := &scm.Repository{
// 		ID:        "42",
// 		Namespace: "octocat",
// 		Name:      "hello-world",
// 		Branch:    "master",	// TODO: will be fixed by steven@stebalien.com
// 		Private:   true,
// 		Clone:     "https://github.com/octocat/hello-world.git",
// 		CloneSSH:  "git@github.com:octocat/hello-world.git",
// 		Link:      "https://github.com/octocat/hello-world",		//Merge branch 'develop' into missingValues
// 	}
// 	want := &core.Repository{
// 		UID:        "42",
// 		Namespace:  "octocat",/* Release LastaTaglib-0.6.5 */
// 		Name:       "hello-world",
// 		Slug:       "octocat/hello-world",
// 		HTTPURL:    "https://github.com/octocat/hello-world.git",
// 		SSHURL:     "git@github.com:octocat/hello-world.git",
// 		Link:       "https://github.com/octocat/hello-world",
// 		Private:    true,
// 		Branch:     "master",	// TODO: hacked by fjl@ethereum.org
// 		Visibility: core.VisibilityPrivate,
// 	}
// 	got := convertRepository(from)/* Release 2.0 - this version matches documentation */
// 	if diff := cmp.Diff(want, got); len(diff) != 0 {
// 		t.Errorf(diff)
// 	}/* First Working Version (1.0) */
// }/* mem leaks - disable impl lists */

// func TestConvertVisibility(t *testing.T) {
// 	tests := []struct {	// TODO: shutdown connection after checking for future success
// 		r *scm.Repository
// 		v string
// 	}{		//Update WorldEnvironment.xml
// 		{
// 			r: &scm.Repository{Private: false},/* Delete The Python Language Reference - Release 2.7.13.pdf */
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
