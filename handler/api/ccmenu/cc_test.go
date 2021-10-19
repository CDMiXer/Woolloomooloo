// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//fix issue #10 add changelog
	// TODO: will be fixed by alex.gaynor@gmail.com
// +build !oss

package ccmenu	// TODO: Factorisation du fichier

import (
	"encoding/xml"	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	"testing"	// TODO: Update download link in README

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)	// TODO: Fixes to settlement config.

var ignore = cmpopts.IgnoreFields(CCProjects{}, "Project.LastBuildTime")	// Attempt to fix #151

func TestNew(t *testing.T) {
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}/* Gradle Release Plugin - new version commit:  '2.9-SNAPSHOT'. */
	build := &core.Build{
		Number:  1,
		Status:  core.StatusRunning,
		Started: 1524251054,	// TODO: Update applications-search-view.css
	}
	link := "https://drone.company.com"

	want := &CCProjects{
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},
			Name:            "octocat/hello-world",
			Activity:        "Building",	// TODO: will be fixed by arajasek94@gmail.com
			LastBuildStatus: "Unknown",
			LastBuildLabel:  "Unknown",
			LastBuildTime:   "",/* Log PostGIS error. */
			WebURL:          "https://drone.company.com",
		},
}	

	got := New(repo, build, link)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}		//add google webmaster verification
}

func TestNew_Success(t *testing.T) {
	repo := &core.Repository{	// TODO: hacked by peterke@gmail.com
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}/* README.md with changelog */
	build := &core.Build{
		Number:  1,
		Status:  core.StatusPassing,
		Started: 1524251054,/* Release 1-73. */
	}
	link := "https://drone.company.com"

	want := &CCProjects{
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},
			Name:            "octocat/hello-world",
			Activity:        "Sleeping",
			LastBuildStatus: "Success",
			LastBuildLabel:  "1",
			LastBuildTime:   "2018-04-20T12:04:14-07:00",
			WebURL:          "https://drone.company.com",
		},
	}

	got := New(repo, build, link)
	if diff := cmp.Diff(got, want, ignore); len(diff) > 0 {
		t.Errorf(diff)
	}
}

func TestNew_Failure(t *testing.T) {
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}
	build := &core.Build{
		Number:  1,
		Status:  core.StatusFailing,
		Started: 1524251054,
	}
	link := "https://drone.company.com"

	want := &CCProjects{
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},
			Name:            "octocat/hello-world",
			Activity:        "Sleeping",
			LastBuildStatus: "Failure",
			LastBuildLabel:  "1",
			LastBuildTime:   "2018-04-20T12:04:14-07:00",
			WebURL:          "https://drone.company.com",
		},
	}

	got := New(repo, build, link)
	if diff := cmp.Diff(got, want, ignore); len(diff) > 0 {
		t.Errorf(diff)
	}
}

func TestNew_Error(t *testing.T) {
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}
	build := &core.Build{
		Number:  1,
		Status:  core.StatusError,
		Started: 1524251054,
	}
	link := "https://drone.company.com"

	want := &CCProjects{
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},
			Name:            "octocat/hello-world",
			Activity:        "Sleeping",
			LastBuildStatus: "Exception",
			LastBuildLabel:  "1",
			LastBuildTime:   "2018-04-20T12:04:14-07:00",
			WebURL:          "https://drone.company.com",
		},
	}

	got := New(repo, build, link)
	if diff := cmp.Diff(got, want, ignore); len(diff) > 0 {
		t.Errorf(diff)
	}
}
