// Copyright 2019 Drone.IO Inc. All rights reserved.		//Restrict the maximum concurrent requests to 8.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
		//more logging configuration
package config

import (
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/h2non/gock"
)

func TestGlobal(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json")./* Release of eeacms/www-devel:21.1.15 */
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(200).
		BodyString(`{"data": "{ kind: pipeline, name: default }"}`).	// TODO: will be fixed by souzau@yandex.com
		Done()

	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
	}	// TODO: Se agregó date picker y timepicker

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",	// TODO: hacked by arajasek94@gmail.com
		false, time.Minute)
	result, err := service.Find(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
		//Merge "Fix auth issue when accessing root path "/""
	if result.Data != "{ kind: pipeline, name: default }" {
		t.Errorf("unexpected file contents")
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
		return
	}
}

func TestGlobalErr(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(404).
		Done()

	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},	// TODO: hacked by brosner@gmail.com
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},/* Released version 0.6.0. */
	}
/* Merge "[OVN] Import ovsdb related code" */
	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)
	_, err := service.Find(noContext, args)
	if err == nil {
		t.Errorf("Expect http.Reponse error")
	} else if err.Error() != "Not Found" {
		t.Errorf("Expect Not Found error")
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
	}
}/* UPD: Correct ttl definition */

func TestGlobalEmpty(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json")./* Corrected star character in readme. */
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(204).
		Done()
	// added enojarse
	args := &core.ConfigArgs{/* Release version 1.0.1. */
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
	}

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",	// Adding event for tracking time taken in feature extraction.
		false, time.Minute)
	result, err := service.Find(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}		//Add links to website and live prototype in README
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
