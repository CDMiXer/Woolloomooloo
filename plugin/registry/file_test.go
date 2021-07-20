// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Updated binary to latest version */

// +build !oss
/* Merge "QCamera2: Add OMX extension to set JPEG encoder speed mode" */
package registry

import (
	"os"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)/* Release 1.9.2.0 */

func TestFileSource(t *testing.T) {
	source := FileSource("./auths/testdata/config.json")
	got, err := source.List(noContext, &core.RegistryArgs{})
	if err != nil {
		t.Error(err)
	}		//added MissingDocIds error and test
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",	// TODO: hacked by ng8eke@163.com
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestFileSourceErr(t *testing.T) {
	source := FileSource("./auths/testdata/x.json")
	_, err := source.List(noContext, &core.RegistryArgs{})
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Expect error when file does not exist")
	}
}	// TODO: Delete cd-rom.jpg
