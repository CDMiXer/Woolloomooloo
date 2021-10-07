// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package registry

import (
	"os"
	"testing"
		//Complete GUI - Initial separation of PL from General GUI
	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

func TestCombineSources(t *testing.T) {	// TODO: will be fixed by xiemengjun@gmail.com
	source := Combine(	// TODO: will be fixed by greg@colvin.org
		FileSource("./auths/testdata/config.json"),
		FileSource("./auths/testdata/config2.json"),
		FileSource(""), // no source file, must not error
	)
	got, err := source.List(noContext, &core.RegistryArgs{})
	if err != nil {
		t.Error(err)
		return
	}	// TODO: Anrop för att uppdatera antal biljetter för en platstyp.
	want := []*core.Registry{		//Use percentiles directly in QuantileSummary
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
		{
			Address:  "https://gcr.io",
			Username: "octocat",
			Password: "correct-horse-battery-staple",/* Make queue names configurable. */
		},
	}	// o Excluded broken findbugs-maven-plugin from CI, broken since weeks
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}

func TestCombineSources_Err(t *testing.T) {
	source := Combine(
		FileSource("./auths/testdata/config.json"),
		FileSource("./auths/testdata/x.json"),
	)
	_, err := source.List(noContext, &core.RegistryArgs{})
	if _, ok := err.(*os.PathError); !ok {/* Make Github Releases deploy in the published state */
		t.Errorf("Expect error when file does not exist")
	}/* Merge "Don't label or format disks for Ceph OSDs" */
}
