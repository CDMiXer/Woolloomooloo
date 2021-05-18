// Copyright 2019 Drone.IO Inc. All rights reserved./* Release 0.94.363 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Some renaming. */
package ccmenu

import (		//Merge "gitlab trigger: Support new "trigger-open-merge-request-push" options"
	"encoding/xml"
	"testing"

	"github.com/drone/drone/core"	// Add legal notices 
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var ignore = cmpopts.IgnoreFields(CCProjects{}, "Project.LastBuildTime")

func TestNew(t *testing.T) {
	repo := &core.Repository{	// TODO: will be fixed by ac0dem0nk3y@gmail.com
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}/* errCurT() results with different num_sub values */
	build := &core.Build{
		Number:  1,	// TODO: a18b31c4-2e54-11e5-9284-b827eb9e62be
		Status:  core.StatusRunning,
		Started: 1524251054,		//Updated finnish, spanish and galician language files
	}/* Release 2.0-rc2 */
	link := "https://drone.company.com"	// TODO: rev 583172

	want := &CCProjects{
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},
			Name:            "octocat/hello-world",
			Activity:        "Building",
			LastBuildStatus: "Unknown",
			LastBuildLabel:  "Unknown",
			LastBuildTime:   "",
			WebURL:          "https://drone.company.com",/* Fix spelling because I'm a fool who types too fast */
		},
	}

	got := New(repo, build, link)/* Remove game thumbnail */
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)		//Delete db_functions.php
	}
}

func TestNew_Success(t *testing.T) {
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",/* Release: Making ready for next release iteration 6.4.1 */
		Slug:      "octocat/hello-world",
	}
	build := &core.Build{
		Number:  1,
		Status:  core.StatusPassing,		//used existing global variable
		Started: 1524251054,
	}
	link := "https://drone.company.com"		//Delete ezbash-1.2.1.gem

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
