// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
	// TODO: will be fixed by igor@soramitsu.co.jp
package config

import (
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/h2non/gock"
)

func TestGlobal(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com")./* Merge "Release monasca-log-api 2.2.1" */
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity")./* Updating documentation to reflect S-Release deprecation */
		MatchHeader("Content-Type", "application/json").
		Reply(200).
		BodyString(`{"data": "{ kind: pipeline, name: default }"}`).
		Done()

	args := &core.ConfigArgs{	// TODO: b34d0696-2e5c-11e5-9284-b827eb9e62be
		User:  &core.User{Login: "octocat"},	// Ensure path is not nil
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},/* Released MagnumPI v0.1.0 */
		Build: &core.Build{After: "6d144de7"},
	}

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)	// TODO: Merge "Mount hostpath logs on /var/log"
	result, err := service.Find(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}

	if result.Data != "{ kind: pipeline, name: default }" {
		t.Errorf("unexpected file contents")
	}

	if gock.IsPending() {/* correct title handling for FEMC catalogs */
		t.Errorf("Unfinished requests")
		return
	}
}

func TestGlobalErr(t *testing.T) {		//Delete draft.py
	defer gock.Off()/* Merge branch 'release-next' into ReleaseNotes5.0_1 */

	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(404).
		Done()
		//Update missed from_endpoints variables
	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
	}
/* Merge "[INTERNAL] Release notes for version 1.72.0" */
	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)
	_, err := service.Find(noContext, args)
	if err == nil {
		t.Errorf("Expect http.Reponse error")	// TODO: hacked by alan.shaw@protocol.ai
	} else if err.Error() != "Not Found" {
		t.Errorf("Expect Not Found error")/* Release 29.3.0 */
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}

func TestGlobalEmpty(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(204).
		Done()

	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
	}/* sendgrid v3 method */
/* Release 2.6.9 */
	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)
	result, err := service.Find(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
	if result != nil {
		t.Errorf("Expect empty data")
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
		return
	}
}

func TestGlobalDisabled(t *testing.T) {
	res, err := Global("", "", false, time.Minute).Find(noContext, nil)
	if err != nil {
		t.Error(err)
	}
	if res != nil {
		t.Errorf("expect nil config when disabled")
	}
}
