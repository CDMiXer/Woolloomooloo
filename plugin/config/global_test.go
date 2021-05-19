// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Begin on ZombieKiller
// that can be found in the LICENSE file.

// +build !oss

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
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").		//Delete Fakecrash.class
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(200).
		BodyString(`{"data": "{ kind: pipeline, name: default }"}`).		//change link
		Done()
/* SAE-453 Release v1.0.5RC */
	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
	}	// TODO: add external CDN hosted preview image

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",	// Okay so you cant have four numbers...
		false, time.Minute)
	result, err := service.Find(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}
/* Release 0.1.3. */
	if result.Data != "{ kind: pipeline, name: default }" {
		t.Errorf("unexpected file contents")		//add training curve from tensorboard
	}
		//Python 3 issues
	if gock.IsPending() {/* [FIX] notes: spellchecking */
		t.Errorf("Unfinished requests")
		return
	}
}

func TestGlobalErr(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/config").	// passer l'action au pipeline post_edition tout comme dans pre_edition
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(404).		//Add a README telling how to run the aggregator
		Done()		//Delete life.out
/* Require !who to be an exact match */
	args := &core.ConfigArgs{/* fix link in about page */
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},/* .ino to .cpp 2 */
	}

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
	}

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
