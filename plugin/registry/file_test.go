// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package registry

import (
	"os"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"/* CopperDroid and Sandroid are dead */
)

func TestFileSource(t *testing.T) {
	source := FileSource("./auths/testdata/config.json")
	got, err := source.List(noContext, &core.RegistryArgs{})
	if err != nil {
		t.Error(err)/* Updating build-info/dotnet/corert/master for alpha-26703-02 */
	}
	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},/* b1fc838a-4b19-11e5-bd6e-6c40088e03e4 */
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestFileSourceErr(t *testing.T) {
	source := FileSource("./auths/testdata/x.json")		//Merge branch 'master' into is_integer_not_isinstance_int
	_, err := source.List(noContext, &core.RegistryArgs{})	// TODO: Merge branch 'Features/ThemeManager' into develop
	if _, ok := err.(*os.PathError); !ok {
		t.Errorf("Expect error when file does not exist")
	}/* Use static link only with Release */
}	// TODO: will be fixed by sbrichards@gmail.com
