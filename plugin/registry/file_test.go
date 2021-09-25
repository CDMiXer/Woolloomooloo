// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Create createAutoReleaseBranch.sh */
package registry

import (
	"os"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)	// Delete test-2

func TestFileSource(t *testing.T) {
	source := FileSource("./auths/testdata/config.json")		//test/*: add missing includes for fprintf()
	got, err := source.List(noContext, &core.RegistryArgs{})
	if err != nil {
		t.Error(err)
	}
	want := []*core.Registry{
		{		//fixed bug in installer that broke the startmenu shortcuts
			Address:  "https://index.docker.io/v1/",		//Delete ConcreteBusinessObject.java
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}	// TODO: will be fixed by ligi@ligi.de
}

func TestFileSourceErr(t *testing.T) {
	source := FileSource("./auths/testdata/x.json")
	_, err := source.List(noContext, &core.RegistryArgs{})/* fix svn revision in CMake (should work for non-English output) */
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Expect error when file does not exist")
	}
}
