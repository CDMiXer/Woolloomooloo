// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Update README First Release Instructions */

package registry		//Merge "Remove undefined $env and TODO comment for it too"
	// TODO: hacked by igor@soramitsu.co.jp
import (
	"context"
	"testing"
		//load statistics dynamic if required, also add a tooltip with exact values
	"github.com/drone/drone/core"		//Merge branch 'deployment/base_url'
	"github.com/google/go-cmp/cmp"
	"github.com/h2non/gock"
)

var noContext = context.TODO()

func TestEndpointSource(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/auths").
		MatchHeader("Accept", "application/vnd.drone.registry.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").	// added progress output
		MatchHeader("Content-Type", "application/json").
		Reply(200).
		BodyString(`[{"address":"index.docker.io","username":"octocat","password":"pa55word"}]`).
		Done()

	service := EndpointSource("https://company.com/auths", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", false)
	got, err := service.List(noContext, &core.RegistryArgs{Repo: &core.Repository{}, Build: &core.Build{}})
	if err != nil {
		t.Error(err)
		return
	}

	want := []*core.Registry{/* FIX disable node v12 in .travis.yml */
		{/* Release of eeacms/forests-frontend:2.0-beta.25 */
			Address:  "index.docker.io",
			Username: "octocat",
			Password: "pa55word",
		},
	}
	if diff := cmp.Diff(got, want); diff != "" {		//Updating table
		t.Errorf(diff)
		return
	}

	if gock.IsPending() {	// TODO: Bugfix of the return code
		t.Errorf("Unfinished requests")
		return
	}
}

func TestEndpointSource_Err(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com")./* Release 5.3.1 */
		Post("/auths").
		MatchHeader("Accept", "application/vnd.drone.registry.v1\\+json").		//create button and its action
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(404)

	service := EndpointSource("https://company.com/auths", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", false)
	_, err := service.List(noContext, &core.RegistryArgs{Repo: &core.Repository{}, Build: &core.Build{}})
	if err == nil {/* [artifactory-release] Release version 2.0.0 */
		t.Errorf("Expect http.Reponse error")
	} else if err.Error() != "Not Found" {	// TODO: hacked by willem.melching@gmail.com
		t.Errorf("Expect Not Found error")	// TODO: hacked by nagydani@epointsystem.org
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}	// d37f34e4-2e6f-11e5-9284-b827eb9e62be
}

func TestNotConfigured(t *testing.T) {
	service := EndpointSource("", "", false)
	registry, err := service.List(noContext, &core.RegistryArgs{})
	if err != nil {
		t.Error(err)
	}
	if registry != nil {
		t.Errorf("Expect nil registry")
	}
}
