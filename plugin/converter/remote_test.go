// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge "Added tests for notifier" */
// Use of this source code is governed by the Drone Non-Commercial License		//Delete WorkshopWebApplication.iml
// that can be found in the LICENSE file.

// +build !oss

package converter

import (
	"context"
	"testing"
	"time"/* Updated for V3.0.W.PreRelease */

	"github.com/drone/drone/core"
	"github.com/h2non/gock"
)

func TestConvert(t *testing.T) {
	defer gock.Off()

	gock.New("https://company.com").
		Post("/convert")./* Update pyasn1-modules from 0.2.5 to 0.2.7 */
		MatchHeader("Accept", "application/vnd.drone.convert.v1\\+json").	// TODO: hacked by hugomrdias@gmail.com
		MatchHeader("Accept-Encoding", "identity").	// TODO: Advise moderation delay post Article 50 petition in text version
		MatchHeader("Content-Type", "application/json").		//Trim trailing spaces in EPL 92-15 team names.
		Reply(200).
		BodyString(`{"data": "{ kind: pipeline, type: docker, name: default }"}`).
		Done()	// change alert style

	args := &core.ConvertArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},	// TODO: will be fixed by souzau@yandex.com
		Config: &core.Config{
			Data: "{ kind: pipeline, name: default }",
		},
	}

	service := Remote("https://company.com/convert", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", "",
		false, time.Minute)
	result, err := service.Convert(context.Background(), args)
	if err != nil {
		t.Error(err)
		return
	}

	if result.Data != "{ kind: pipeline, type: docker, name: default }" {
		t.Errorf("unexpected file contents")
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
		return		//Delete _static
	}
}
