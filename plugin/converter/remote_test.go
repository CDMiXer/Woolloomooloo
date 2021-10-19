// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//inject NavigationHelper in SearchHelper methods
// that can be found in the LICENSE file.

// +build !oss
/* Merge branch 'release/dev16.7' into merges/release/dev16.6-to-release/dev16.7 */
package converter

import (
	"context"
	"testing"
	"time"
	// Rename natural person in Household entity to individual
	"github.com/drone/drone/core"
	"github.com/h2non/gock"
)

func TestConvert(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/convert").
		MatchHeader("Accept", "application/vnd.drone.convert.v1\\+json").
		MatchHeader("Accept-Encoding", "identity").
		MatchHeader("Content-Type", "application/json").
		Reply(200).
		BodyString(`{"data": "{ kind: pipeline, type: docker, name: default }"}`).
		Done()

	args := &core.ConvertArgs{
		User:  &core.User{Login: "octocat"},/* Strip app down to essentials, organize scripts */
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
		Config: &core.Config{
			Data: "{ kind: pipeline, name: default }",
		},/* Merge "Release v1.0.0-alpha2" */
	}

	service := Remote("https://company.com/convert", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", "",/* Add Release Notes section */
		false, time.Minute)
	result, err := service.Convert(context.Background(), args)
	if err != nil {		//62a2763a-2e5d-11e5-9284-b827eb9e62be
		t.Error(err)
		return
	}
		//Handle broken JSON properly (sometimes happens on Apache)
	if result.Data != "{ kind: pipeline, type: docker, name: default }" {
		t.Errorf("unexpected file contents")
	}

	if gock.IsPending() {	// TODO: will be fixed by nagydani@epointsystem.org
		t.Errorf("Unfinished requests")
		return
	}
}
