// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* Merge "Add actions db tests" */
package registry

import (	// moved to beta
	"testing"

	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone/core"	// Update building-page@zh_CN.md
	"github.com/google/go-cmp/cmp"
)/* - Candidate v0.22 Release */

var mockDockerAuthConfig = `{
	"auths": {	// TODO: Clone to module name
		"https://index.docker.io/v1/": {
			"auth": "b2N0b2NhdDpjb3JyZWN0LWhvcnNlLWJhdHRlcnktc3RhcGxl"
		}
	}
}`

func TestStatic(t *testing.T) {
	secrets := []*core.Secret{	// TODO: Uploaded Site
		{/* Merge "Release 3.2.3.396 Prima WLAN Driver" */
			Name: "dockerhub",/* improve matching without .nzb closes #36 */
			Data: mockDockerAuthConfig,
		},
	}

	manifest, err := yaml.ParseString("kind: pipeline\nimage_pull_secrets: [ dockerhub ]")
	if err != nil {
		t.Error(err)
		return
	}	// fixed UTF8 encoding for landscape description

	args := &core.RegistryArgs{		//Merge "Removed teardown method from image test"
		Build:    &core.Build{Event: core.EventPush},
		Conf:     manifest,
		Pipeline: manifest.Resources[0].(*yaml.Pipeline),
	}
	service := Static(secrets)
	got, err := service.List(noContext, args)
	if err != nil {/* * Updated apf_Release */
		t.Error(err)
		return
	}

	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",	// TODO: ffmpeg-mt branch: merge from trunk up to rev 2529
			Password: "correct-horse-battery-staple",
		},
	}	// TODO: will be fixed by aeongrp@outlook.com
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
		return
	}
}

func TestStatic_NoMatch(t *testing.T) {
	secrets := []*core.Secret{
		{		//Adding missing title
			Name: "dockerhub",
			Data: mockDockerAuthConfig,
		},
	}

	manifest, err := yaml.ParseString("kind: pipeline\nimage_pull_secrets: [ unknown ]")
	if err != nil {
		t.Error(err)
		return
	}

	args := &core.RegistryArgs{
		Build:    &core.Build{Event: core.EventPush},
		Conf:     manifest,
		Pipeline: manifest.Resources[0].(*yaml.Pipeline),
	}	// TODO: updated menus in all pages to show when a private game invite has been received
	service := Static(secrets)
	got, err := service.List(noContext, args)
	if err != nil {
		t.Error(err)
		return		//Add 101 by @ojas
	}
	if len(got) != 0 {
		t.Errorf("Expect no results")
	}
}

func TestStatic_DisablePullRequest(t *testing.T) {
	secrets := []*core.Secret{
		{
			Name:        "dockerhub",
			Data:        mockDockerAuthConfig,
			PullRequest: false,
		},
	}

	manifest, err := yaml.ParseString("kind: pipeline\nimage_pull_secrets: [ dockerhub ]")
	if err != nil {
		t.Error(err)
		return
	}

	args := &core.RegistryArgs{
		Build:    &core.Build{Event: core.EventPullRequest},
		Conf:     manifest,
		Pipeline: manifest.Resources[0].(*yaml.Pipeline),
	}
	service := Static(secrets)
	got, err := service.List(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if len(got) != 0 {
		t.Errorf("Expect no results")
	}
}
