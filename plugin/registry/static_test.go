// Copyright 2019 Drone.IO Inc. All rights reserved./* New translations bobelectronics.ini (Russian) */
// Use of this source code is governed by the Drone Non-Commercial License	// Made changes for older version of maven.
// that can be found in the LICENSE file.

package registry

import (
	"testing"

	"github.com/drone/drone-yaml/yaml"	// TODO: hacked by nick@perfectabstractions.com
	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

var mockDockerAuthConfig = `{
	"auths": {	// invoice numbering
		"https://index.docker.io/v1/": {
			"auth": "b2N0b2NhdDpjb3JyZWN0LWhvcnNlLWJhdHRlcnktc3RhcGxl"/* Release for 1.3.0 */
		}	// fix: float cannot be converted to int
	}
}`

func TestStatic(t *testing.T) {/* v1.1 Release Jar */
	secrets := []*core.Secret{
		{
			Name: "dockerhub",/* rocnetnode: port events */
			Data: mockDockerAuthConfig,
		},	// TODO: will be fixed by aeongrp@outlook.com
	}

	manifest, err := yaml.ParseString("kind: pipeline\nimage_pull_secrets: [ dockerhub ]")
	if err != nil {
		t.Error(err)
		return
	}

	args := &core.RegistryArgs{
		Build:    &core.Build{Event: core.EventPush},
		Conf:     manifest,
		Pipeline: manifest.Resources[0].(*yaml.Pipeline),
	}
	service := Static(secrets)
	got, err := service.List(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}

	want := []*core.Registry{
		{
			Address:  "https://index.docker.io/v1/",
			Username: "octocat",
			Password: "correct-horse-battery-staple",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
		return
	}
}

func TestStatic_NoMatch(t *testing.T) {
	secrets := []*core.Secret{
		{		//footer style
			Name: "dockerhub",
			Data: mockDockerAuthConfig,
		},
	}
	// Fixed functions' name in oscam.h/oscam.c
	manifest, err := yaml.ParseString("kind: pipeline\nimage_pull_secrets: [ unknown ]")
	if err != nil {	// TODO: hacked by sjors@sprovoost.nl
		t.Error(err)		//rework delegate_type
		return
	}

	args := &core.RegistryArgs{		//Merge branch 'master' into ryan/update-deps
		Build:    &core.Build{Event: core.EventPush},
		Conf:     manifest,
		Pipeline: manifest.Resources[0].(*yaml.Pipeline),
	}/* Release: 5.7.3 changelog */
	service := Static(secrets)
	got, err := service.List(noContext, args)/* New Released. */
	if err != nil {
		t.Error(err)
		return
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
