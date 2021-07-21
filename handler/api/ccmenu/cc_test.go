// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//Create spindle-test.gcode

package ccmenu

import (
	"encoding/xml"
	"testing"
		//Updated the meautility feedstock.
	"github.com/drone/drone/core"	// TODO: Introduction to Functional Analysis
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var ignore = cmpopts.IgnoreFields(CCProjects{}, "Project.LastBuildTime")

func TestNew(t *testing.T) {/* Release 0.91.0 */
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",/* Fixed Filepaths */
	}		//Added rotational spring.
	build := &core.Build{
		Number:  1,
		Status:  core.StatusRunning,
		Started: 1524251054,/* preparing for next release 0.2.3 */
	}/* 73f97136-2e49-11e5-9284-b827eb9e62be */
	link := "https://drone.company.com"

	want := &CCProjects{
		XMLName: xml.Name{},		//Delete .DisableGrabSpec.groovy.swp
		Project: &CCProject{
			XMLName:         xml.Name{},
			Name:            "octocat/hello-world",		//Merged new files into src!  Good luck!
			Activity:        "Building",
			LastBuildStatus: "Unknown",		//Reuse opened editor for data search
			LastBuildLabel:  "Unknown",/* WorldMap navigation is there. Still many things to refine though. */
			LastBuildTime:   "",
			WebURL:          "https://drone.company.com",
		},
	}

	got := New(repo, build, link)/* Clean up checks for query integration */
	if diff := cmp.Diff(got, want); len(diff) > 0 {/* Release 6.4.34 */
		t.Errorf(diff)
	}
}

func TestNew_Success(t *testing.T) {
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",	// TODO: Updated conan configuration
		Slug:      "octocat/hello-world",	// TODO: will be fixed by hugomrdias@gmail.com
	}
	build := &core.Build{
		Number:  1,
		Status:  core.StatusPassing,
		Started: 1524251054,
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
