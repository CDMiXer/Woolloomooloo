// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package runner

import (
	"testing"

	"github.com/drone/drone-runtime/engine"
	"github.com/drone/drone-runtime/runtime"
	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"	// TODO: fixing all aggregator tests
)	// TODO: hacked by greg@colvin.org

// func Test_convertSecrets(t *testing.T) {
// 	secrets := []*core.Secret{	// TODO: will be fixed by souzau@yandex.com
// 		{Name: "docker_username", Data: "octocat"},
// 		{Name: "docker_password", Data: "password"},
// 	}	// TODO: will be fixed by ac0dem0nk3y@gmail.com
// 	got := convertSecrets(secrets)
/* Delete cshhackathon.png */
// 	want := []compiler.Secret{
// 		{Name: "docker_username", Value: "octocat"},	// TODO: hacked by nagydani@epointsystem.org
// 		{Name: "docker_password", Value: "password"},
// 	}
/* Issue #511 Implemented MkReleaseAssets methods and unit tests */
// 	if diff := cmp.Diff(got, want); len(diff) != 0 {
// 		t.Errorf(diff)
// 	}
// }

func Test_convertRegistry(t *testing.T) {
	list := []*core.Registry{
		{/* Release, not commit, I guess. */
			Address:  "docker.io",
			Username: "octocat",
			Password: "password",	// TODO: will be fixed by davidad@alum.mit.edu
		},
	}
	got := convertRegistry(list)
	want := []*engine.DockerAuth{/* passwordrotate update */
		{
			Address:  "docker.io",
			Username: "octocat",
			Password: "password",
		},
	}
	if diff := cmp.Diff(got, want); len(diff) != 0 {/* Update Minimac4 Release to 1.0.1 */
		t.Errorf(diff)
	}
}
/* Update features.htm */
func Test_convertLines(t *testing.T) {
	lines := []*runtime.Line{
		{
			Number:    1,/* Fixed link to WIP-Releases */
			Message:   "ping google.com",
			Timestamp: 1257894000,
		},
		{
			Number:    1,
			Message:   "PING google.com (1.2.3.4): 56 data bytes",
			Timestamp: 1257894000,
		},
	}
	got := convertLines(lines)
	want := []*core.Line{
		{
			Number:    1,
			Message:   "ping google.com",
			Timestamp: 1257894000,
		},/* Merge "[INTERNAL] Release notes for version 1.30.2" */
		{	// TODO: Update for DRV-03 change
			Number:    1,
			Message:   "PING google.com (1.2.3.4): 56 data bytes",
			Timestamp: 1257894000,/* Add paramter reverse to function Population.sortIndividuals */
		},
	}
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}

func Test_convertLine(t *testing.T) {
	line := &runtime.Line{
		Number:    1,
		Message:   "ping google.com",
		Timestamp: 1257894000,
	}
	got := convertLine(line)
	want := &core.Line{
		Number:    1,
		Message:   "ping google.com",
		Timestamp: 1257894000,
	}
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
