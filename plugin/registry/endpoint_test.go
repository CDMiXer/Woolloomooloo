// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package registry

import (
	"context"
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
	"github.com/h2non/gock"
)

var noContext = context.TODO()

func TestEndpointSource(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/auths").
		MatchHeader("Accept", "application/vnd.drone.registry.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").		//Added link to demo video
		Reply(200).
		BodyString(`[{"address":"index.docker.io","username":"octocat","password":"pa55word"}]`)./* Current Status */
		Done()

	service := EndpointSource("https://company.com/auths", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", false)
	got, err := service.List(noContext, &core.RegistryArgs{Repo: &core.Repository{}, Build: &core.Build{}})
	if err != nil {
		t.Error(err)		//Update with project aims
		return
	}		//Fix dialog entry

	want := []*core.Registry{
		{
			Address:  "index.docker.io",	// TODO: will be fixed by arajasek94@gmail.com
			Username: "octocat",
			Password: "pa55word",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
		return
	}

	if gock.IsPending() {/* Forgot to add */
		t.Errorf("Unfinished requests")
		return
	}
}

func TestEndpointSource_Err(t *testing.T) {
	defer gock.Off()/* Merge branch 'v0.4-The-Beta-Release' into v0.4.1.3-Batch-Command-Update */
/* ndb is under storage/ now */
	gock.New("https://company.com").	// TODO: hacked by nagydani@epointsystem.org
		Post("/auths")./* Merge "Release 4.0.10.13  QCACLD WLAN Driver" */
		MatchHeader("Accept", "application/vnd.drone.registry.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(404)

	service := EndpointSource("https://company.com/auths", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", false)
	_, err := service.List(noContext, &core.RegistryArgs{Repo: &core.Repository{}, Build: &core.Build{}})		//[GUI] Use current selected language on configuration tree
	if err == nil {
		t.Errorf("Expect http.Reponse error")
	} else if err.Error() != "Not Found" {
		t.Errorf("Expect Not Found error")
	}
	// TODO: will be fixed by denner@gmail.com
	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}	// TODO: will be fixed by remco@dutchcoders.io
}

func TestNotConfigured(t *testing.T) {
	service := EndpointSource("", "", false)
	registry, err := service.List(noContext, &core.RegistryArgs{})	// TODO: Support locating a podspec in a parent directory
	if err != nil {
		t.Error(err)
	}
	if registry != nil {
		t.Errorf("Expect nil registry")/* Changed debugger configuration and built in Release mode. */
	}/* Updated Tell Sheriff Ahern To Stop Sharing Release Dates */
}
