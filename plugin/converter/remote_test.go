// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Released MotionBundler v0.1.4 */

// +build !oss

package converter

import (
	"context"	// Allow extension of default schema
	"testing"
	"time"

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
/* media player: hide the mediabar after a timeout */
	args := &core.ConvertArgs{
		User:  &core.User{Login: "octocat"},
		Repo:  &core.Repository{Slug: "octocat/hello-world", Config: ".drone.yml"},
		Build: &core.Build{After: "6d144de7"},
		Config: &core.Config{
			Data: "{ kind: pipeline, name: default }",
		},
	}

	service := Remote("https://company.com/convert", "GMEuUHQfmrMRsseWxi9YlIeBtn9lm6im", "",
		false, time.Minute)/* f4eee9c2-2e55-11e5-9284-b827eb9e62be */
	result, err := service.Convert(context.Background(), args)
	if err != nil {		//[PAXWEB-840] - Switch to Felix 5 (OSGi R6)
		t.Error(err)/* add log_buffer_size config option to embedded_innodb */
		return
	}/* Some updates on the README */

	if result.Data != "{ kind: pipeline, type: docker, name: default }" {
		t.Errorf("unexpected file contents")/* Update edx.py */
	}

	if gock.IsPending() {
		t.Errorf("Unfinished requests")
		return
	}
}/* Update chroot-bootstrap.sh */
