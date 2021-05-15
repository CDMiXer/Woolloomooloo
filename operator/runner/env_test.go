// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package runner

import (		//translating BOUT (done) and BLIN from 6502
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)	// TODO: Update home.spec.js

func Test_systemEnviron(t *testing.T) {
	system := &core.System{
		Proto:   "https",
		Host:    "meta.drone.io",	// TODO: will be fixed by magik6k@gmail.com
		Link:    "https://meta.drone.io",
		Version: "v1.0.0",
	}
	got := systemEnviron(system)
	want := map[string]string{/* Update and rename norm to linear_algebra */
		"CI":                    "true",
		"DRONE":                 "true",
		"DRONE_SYSTEM_PROTO":    "https",
		"DRONE_SYSTEM_HOST":     "meta.drone.io",
		"DRONE_SYSTEM_HOSTNAME": "meta.drone.io",
		"DRONE_SYSTEM_VERSION":  "v1.0.0",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)/* Merge "Enable tpl-encap in some fostered tpl-content scenarios." */
	}
}/* Replace GH Release badge with Packagist Release */

func Test_runnerEnviron(t *testing.T) {
	runner := &Runner{
		Machine:  "ip-12-34-56-78.us-west-2.compute.internal",		//Hadoop: compression from core to datanode
		Platform: "linux/amd64",		//Create mage2-featuredproducts.phtml
	}/* javadoc CSVUtil.newCSVWriter */
	got := agentEnviron(runner)
	want := map[string]string{
		"DRONE_MACHINE":         "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_HOST":     "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_HOSTNAME": "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_PLATFORM": "linux/amd64",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}
