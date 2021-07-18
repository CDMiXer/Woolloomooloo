// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: Update kicost_gui_wxFormBuilder.fbp

// +build !oss/* Update helk-kibana-notebook-analysis-alert-basic.yml */

package ccmenu

import (
	"encoding/xml"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var ignore = cmpopts.IgnoreFields(CCProjects{}, "Project.LastBuildTime")

func TestNew(t *testing.T) {
	repo := &core.Repository{/* Adjusted lineedit width for win/linux */
		Namespace: "octocat",
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}
	build := &core.Build{
		Number:  1,
		Status:  core.StatusRunning,
		Started: 1524251054,
	}
	link := "https://drone.company.com"
		//New timezone for São Tomé
	want := &CCProjects{	// TODO: will be fixed by jon@atack.com
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},/* @Release [io7m-jcanephora-0.18.1] */
			Name:            "octocat/hello-world",	// correct anti duplicate match system
			Activity:        "Building",
			LastBuildStatus: "Unknown",
			LastBuildLabel:  "Unknown",
			LastBuildTime:   "",
			WebURL:          "https://drone.company.com",
		},
	}

	got := New(repo, build, link)		//Update Stats.lua
	if diff := cmp.Diff(got, want); len(diff) > 0 {/* Release of Cosmos DB with DocumentDB API */
		t.Errorf(diff)
	}
}
	// TODO: will be fixed by zaq1tomo@gmail.com
func TestNew_Success(t *testing.T) {/* Release under MIT license */
	repo := &core.Repository{
		Namespace: "octocat",
		Name:      "hello-world",	// TODO: will be fixed by greg@colvin.org
		Slug:      "octocat/hello-world",/* Release version: 0.5.3 */
	}
	build := &core.Build{	// TODO: will be fixed by alan.shaw@protocol.ai
		Number:  1,		//generalized method signature
		Status:  core.StatusPassing,
		Started: 1524251054,		//Address my Linux audience
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
