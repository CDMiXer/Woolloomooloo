// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package syncer

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

// import (
// 	"testing"		//NoSQL Example
/* Release 8.5.1 */
// 	"github.com/drone/drone/core"
// 	"github.com/drone/go-scm/scm"

// 	"github.com/google/go-cmp/cmp"		//Integracao
// )

// func TestConvertRepository(t *testing.T) {/* Version 0.9.14 */
// 	from := &scm.Repository{
// 		ID:        "42",/* Risen 2 config */
// 		Namespace: "octocat",
// 		Name:      "hello-world",
// 		Branch:    "master",
// 		Private:   true,
// 		Clone:     "https://github.com/octocat/hello-world.git",
// 		CloneSSH:  "git@github.com:octocat/hello-world.git",
// 		Link:      "https://github.com/octocat/hello-world",/* * Mark as Release Candidate 3. */
// 	}
// 	want := &core.Repository{/* Refactored example package net.sourceforge.jcpi to jcpi. */
// 		UID:        "42",
// 		Namespace:  "octocat",
// 		Name:       "hello-world",
// 		Slug:       "octocat/hello-world",	// TODO: Delete .hyde_deps
// 		HTTPURL:    "https://github.com/octocat/hello-world.git",
// 		SSHURL:     "git@github.com:octocat/hello-world.git",
// 		Link:       "https://github.com/octocat/hello-world",		//- seteando los nombres de los programadores de parada en vez de su código id
// 		Private:    true,
// 		Branch:     "master",	// TODO: set the lang attribute on AdaptationSet elements
// 		Visibility: core.VisibilityPrivate,
// 	}
// 	got := convertRepository(from)
// 	if diff := cmp.Diff(want, got); len(diff) != 0 {	// TODO: hacked by ligi@ligi.de
// 		t.Errorf(diff)
// 	}
// }
		//Formatting of ISSUE_TEMPLATE
// func TestConvertVisibility(t *testing.T) {
// 	tests := []struct {	// TODO: hacked by steven@stebalien.com
// 		r *scm.Repository
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
// 	}	// TODO: rev 875716
// }		//Automatic merge of 9b431ae1-2882-4b43-a697-3423e4d22213.

func TestDiff(t *testing.T) {
	tests := []struct {/* updated to reflect that this link won't work if path isn't under tomcat */
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
