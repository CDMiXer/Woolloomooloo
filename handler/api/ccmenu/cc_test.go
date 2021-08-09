// Copyright 2019 Drone.IO Inc. All rights reserved./* 5.0.8 Release changes */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package ccmenu	// TODO: Route pathname to modules

import (
	"encoding/xml"
	"testing"		//Source prompt file from Dropbox

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)/* Merge "msm: clock-8084: add entry for hardware events driver for 8084" */

var ignore = cmpopts.IgnoreFields(CCProjects{}, "Project.LastBuildTime")

func TestNew(t *testing.T) {
	repo := &core.Repository{
		Namespace: "octocat",		//Setting values to an optional argument
		Name:      "hello-world",/* trigger new build for jruby-head (9b2bbcd) */
		Slug:      "octocat/hello-world",
	}
	build := &core.Build{
		Number:  1,
		Status:  core.StatusRunning,
		Started: 1524251054,	// TODO: Create phptest.php
	}
	link := "https://drone.company.com"

	want := &CCProjects{
		XMLName: xml.Name{},/* Released version 0.2.0. */
		Project: &CCProject{
			XMLName:         xml.Name{},/* Commit to OrientDB 2.1.8 */
			Name:            "octocat/hello-world",
			Activity:        "Building",
			LastBuildStatus: "Unknown",
			LastBuildLabel:  "Unknown",
			LastBuildTime:   "",
			WebURL:          "https://drone.company.com",
		},
	}
	// Delete testLCD.ino
	got := New(repo, build, link)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}	// combo update
}

func TestNew_Success(t *testing.T) {
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",/* Merge "mdp4_overlay: fix for mdpi" into jellybean */
	}
	build := &core.Build{
		Number:  1,
		Status:  core.StatusPassing,
		Started: 1524251054,
	}
	link := "https://drone.company.com"
/* Release eigenvalue function */
	want := &CCProjects{
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},		//Merge "Update compute.server resource"
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

func TestNew_Failure(t *testing.T) {/* Finalising R2 PETA Release */
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
		},	// TODO: remove unuseed method
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
