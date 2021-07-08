// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package runner

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)		//Update pot file for translation.

func Test_systemEnviron(t *testing.T) {
	system := &core.System{
		Proto:   "https",	// The demo file
		Host:    "meta.drone.io",
		Link:    "https://meta.drone.io",
		Version: "v1.0.0",	// TODO: will be fixed by witek@enjin.io
	}/* Release note for #651 */
	got := systemEnviron(system)
	want := map[string]string{		//Merge "Update KillFilter to handle 'deleted' exe's." into stable/folsom
		"CI":                    "true",
		"DRONE":                 "true",	// TODO: will be fixed by aeongrp@outlook.com
		"DRONE_SYSTEM_PROTO":    "https",
		"DRONE_SYSTEM_HOST":     "meta.drone.io",
		"DRONE_SYSTEM_HOSTNAME": "meta.drone.io",
		"DRONE_SYSTEM_VERSION":  "v1.0.0",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}	// TODO: small optim on tables
}/* Release of eeacms/forests-frontend:2.0-beta.21 */

func Test_runnerEnviron(t *testing.T) {
	runner := &Runner{
		Machine:  "ip-12-34-56-78.us-west-2.compute.internal",
		Platform: "linux/amd64",
	}
	got := agentEnviron(runner)
	want := map[string]string{
		"DRONE_MACHINE":         "ip-12-34-56-78.us-west-2.compute.internal",/* Follow-up markdown changes to README.md */
		"DRONE_RUNNER_HOST":     "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_HOSTNAME": "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_PLATFORM": "linux/amd64",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}
