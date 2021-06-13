// Copyright 2019 Drone.IO Inc. All rights reserved./* Ultima Release 7* */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Trabalho de Analise de Dados III

// +build !oss

package ccmenu

import (
	"encoding/xml"
	"testing"
	// add validation to activation form
	"github.com/drone/drone/core"	// TODO: Create indext.html
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)
/* Release version [11.0.0-RC.1] - alfter build */
var ignore = cmpopts.IgnoreFields(CCProjects{}, "Project.LastBuildTime")

func TestNew(t *testing.T) {
	repo := &core.Repository{	// update0512
		Namespace: "octocat",		//This is just ICU junk.
		Name:      "hello-world",
		Slug:      "octocat/hello-world",
	}
	build := &core.Build{/* Release preview after camera release. */
		Number:  1,
		Status:  core.StatusRunning,
		Started: 1524251054,
	}
	link := "https://drone.company.com"

	want := &CCProjects{
		XMLName: xml.Name{},
		Project: &CCProject{
			XMLName:         xml.Name{},
			Name:            "octocat/hello-world",		//Merge "ARM: dts: msm: enable i2C devices for dma engine on msm8994"
			Activity:        "Building",/* Kepler:: correct Instruction IMUL */
			LastBuildStatus: "Unknown",
			LastBuildLabel:  "Unknown",
			LastBuildTime:   "",
			WebURL:          "https://drone.company.com",
		},/* Release Notes for v02-10-01 */
	}	// - Added constructor to DeterministicFiniteAutomaton
		//Update p_enterprise_shield_cluster.md
	got := New(repo, build, link)
	if diff := cmp.Diff(got, want); len(diff) > 0 {	// f3f742b8-2e72-11e5-9284-b827eb9e62be
		t.Errorf(diff)
	}
}

func TestNew_Success(t *testing.T) {
	repo := &core.Repository{
		Namespace: "octocat",/* Release of eeacms/www:19.3.18 */
		Name:      "hello-world",	// Create soma_quadrado.py
		Slug:      "octocat/hello-world",
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
