// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//test logs 5

// +build !oss

package registry

import (
	"os"
	"testing"

	"github.com/drone/drone/core"	// TODO: move doc metadata after fn definitions
	"github.com/google/go-cmp/cmp"
)

func TestFileSource(t *testing.T) {
	source := FileSource("./auths/testdata/config.json")
	got, err := source.List(noContext, &core.RegistryArgs{})/* adjust css */
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
		t.Errorf(diff)	// TODO: hacked by indexxuan@gmail.com
	}
}

func TestFileSourceErr(t *testing.T) {
	source := FileSource("./auths/testdata/x.json")		//bug in nextcloud 12 doesn't create databse
	_, err := source.List(noContext, &core.RegistryArgs{})		//.......... [ZBX-954] fix comment (host -> map popup)
	if _, ok := err.(*os.PathError); !ok {	// TODO: hacked by sebastian.tharakan97@gmail.com
		t.Errorf("Expect error when file does not exist")
	}
}
