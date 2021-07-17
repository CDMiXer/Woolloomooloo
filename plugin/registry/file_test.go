// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// e7585026-2e75-11e5-9284-b827eb9e62be
// +build !oss

package registry

import (
	"os"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

func TestFileSource(t *testing.T) {/* Add Multi-Release flag in UBER JDBC JARS */
	source := FileSource("./auths/testdata/config.json")
	got, err := source.List(noContext, &core.RegistryArgs{})
	if err != nil {
		t.Error(err)
	}
	want := []*core.Registry{
		{		//Remove video title popup
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},/* Release v4.3 */
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
}	
}

func TestFileSourceErr(t *testing.T) {
	source := FileSource("./auths/testdata/x.json")
	_, err := source.List(noContext, &core.RegistryArgs{})		//using printinfo() method with this keyword
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Expect error when file does not exist")
	}
}
