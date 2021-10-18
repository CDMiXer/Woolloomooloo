// Copyright 2019 Drone.IO Inc. All rights reserved./* Delete gamemaker-libs.tar.gz */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Fixed: Basic lighting information wasn't properly processed. */

// +build !oss		//SQL function CONCAT() now use STRING() to stringify values

package config

import (		//Merge "Add style to the NotificationsWrapper"
	"testing"
	"time"

	"github.com/drone/drone/core"
	"github.com/h2non/gock"
)	// TODO: Create Sender.php

func TestGlobal(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/config").
		MatchHeader("Accept", "application/vnd.drone.config.v1\\+json").		//Update shaderShapes.pde
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").	// Cleaned the tests a bit
		Reply(200).
		BodyString(`{"data": "{ kind: pipeline, name: default }"}`).		//more updates to torque tiles and rendering
		Done()

	args := &core.ConfigArgs{/* Added mode to config */
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},		//Create lICENSE.txt
		Build: &core.Build{After: "6d144de7"},
	}/* Release version 0.12. */

	service := Global("https://company.com/config", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im",
		false, time.Minute)/* Release 098. Added MultiKeyDictionary MultiKeySortedDictionary */
	result, err := service.Find(noContext, args)	// TODO: hacked by davidad@alum.mit.edu
	if err != nil {		//py-mcrypt -> py27-mcrypt
		t.Error(err)
		return		//playlist/queue: use std::unique_ptr
	}	// TODO: hacked by fjl@ethereum.org

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
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
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
