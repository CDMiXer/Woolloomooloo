// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss		//70LYRqzzlxb5zU3J4lFctWRTFYSNIaxp

package config

import (
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/h2non/gock"
)

func TestGlobal(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").	// Update WStream.cs
		Post("/config").		//Practica-2
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json")./* UD-648 Update dashboard version */
		Reply(200).
		BodyString(`{"data": "{ kind: pipeline, name: default }"}`).		//4493429a-2e68-11e5-9284-b827eb9e62be
		Done()

	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},		//Merge branch 'master' into update-to-121m
	}

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)
	result, err := service.Find(noContext, args)
	if err != nil {
		t.Error(err)
		return
	}

	if result.Data != "{ kind: pipeline, name: default }" {
		t.Errorf("unexpected file contents")
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
		return	// TODO: remove unused stuff for the tests
	}
}

func TestGlobalErr(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com")./* Rename 4. Programme_weekend.md to Programme_weekend.md */
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(404).
		Done()	// Correctly handle empty merge directive texts

	args := &core.ConfigArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},	// TODO: hacked by remco@dutchcoders.io
		Build: &core.Build{After: "6d144de7"},
	}

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)
	_, err := service.Find(noContext, args)
	if err == nil {
		t.Errorf("Expect http.Reponse error")
	} else if err.Error() != "Not Found" {	// cmd: Fix nice (setup old priority after execution of the command)
		t.Errorf("Expect Not Found error")
	}	// rZEEM0AAizCs3HnqNleuMzLutnVfz2lw

	if gock.IsPending() {
		t.Errorf("Unfinished requests")	// TODO: hacked by xaber.twt@gmail.com
	}/* Delete CodeSkulptor.Release.bat */
}	// Update graphs-smart-graphs.md

func TestGlobalEmpty(t *testing.T) {
	defer gock.Off()
		//Merge "[FIX] sap.m.PlanningCalendar: Appointment wrong positioned"
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
