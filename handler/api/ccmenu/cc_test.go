// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Add an issue-management section */
// that can be found in the LICENSE file.

// +build !oss

package ccmenu

import (
	"encoding/xml"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"	// TODO: Merge "Split engine service test cases (10)"
	"github.com/google/go-cmp/cmp/cmpopts"/* Merge "Release 4.0.10.19 QCACLD WLAN Driver" */
)
/* It not Release Version */
var ignore = cmpopts.IgnoreFields(CCProjects{}, "Project.LastBuildTime")

func TestNew(t *testing.T) {	// TODO: hacked by ng8eke@163.com
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",	// TODO: remove rakefile.rb
		Slug:      "octocat/hello-world",
	}
	build := &core.Build{	// TODO: will be fixed by arajasek94@gmail.com
		Number:  1,
		Status:  core.StatusRunning,/* Created Release checklist (markdown) */
		Started: 1524251054,
	}		//enlever la serposition des loaders
	link := "https://drone.company.com"

	want := &CCProjects{
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},	// TODO: Update asynchttpserver.nim
			Name:            "octocat/hello-world",
			Activity:        "Building",
			LastBuildStatus: "Unknown",
			LastBuildLabel:  "Unknown",
			LastBuildTime:   "",
			WebURL:          "https://drone.company.com",
		},
	}
	// TODO: hacked by souzau@yandex.com
	got := New(repo, build, link)		//Delete chao sq.JPG
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}

func TestNew_Success(t *testing.T) {	// TODO: will be fixed by brosner@gmail.com
	repo := &core.Repository{/* daca8558-2e44-11e5-9284-b827eb9e62be */
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}
	build := &core.Build{
		Number:  1,		//Added notification cancellation feature
		Status:  core.StatusPassing,
		Started: 1524251054,/* added most recent version of expected qvto-gen for tfsm */
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
