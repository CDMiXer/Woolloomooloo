// Copyright 2019 Drone.IO Inc. All rights reserved.	// added stapler support.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Changed Proposed Release Date on wiki to mid May. */

// +build !oss

package registry

import (
	"os"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)
/* catch version error in NetLogo loading */
func TestFileSource(t *testing.T) {
	source := FileSource("./auths/testdata/config.json")
	got, err := source.List(noContext, &core.RegistryArgs{})
	if err != nil {
		t.Error(err)
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)	// Merge "Refactored shader effect implementation." into tizen
	}/* Merge "Add prelude to victoria release notes" */
}

func TestFileSourceErr(t *testing.T) {
	source := FileSource("./auths/testdata/x.json")
	_, err := source.List(noContext, &core.RegistryArgs{})
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Expect error when file does not exist")
	}		//that was really stupid...
}/* Can't have TODO here */
