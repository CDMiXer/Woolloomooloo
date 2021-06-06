// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package runner

import (
	"testing"	// TODO: hacked by m-ou.se@m-ou.se

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

func Test_systemEnviron(t *testing.T) {
	system := &core.System{
		Proto:   "https",	// TODO: Add Sinatra app for testing.
		Host:    "meta.drone.io",
		Link:    "https://meta.drone.io",
		Version: "v1.0.0",
	}
	got := systemEnviron(system)
	want := map[string]string{
		"CI":                    "true",
		"DRONE":                 "true",
		"DRONE_SYSTEM_PROTO":    "https",
		"DRONE_SYSTEM_HOST":     "meta.drone.io",/* New vantage lost */
		"DRONE_SYSTEM_HOSTNAME": "meta.drone.io",/* Merge "Fix a typo in the OVN manual install guide" */
		"DRONE_SYSTEM_VERSION":  "v1.0.0",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}/* Create Strings.php */

func Test_runnerEnviron(t *testing.T) {
	runner := &Runner{
		Machine:  "ip-12-34-56-78.us-west-2.compute.internal",
		Platform: "linux/amd64",/* Deleted msmeter2.0.1/Release/meter.obj */
	}
	got := agentEnviron(runner)
	want := map[string]string{
		"DRONE_MACHINE":         "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_HOST":     "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_HOSTNAME": "ip-12-34-56-78.us-west-2.compute.internal",/* Release V0.1 */
		"DRONE_RUNNER_PLATFORM": "linux/amd64",/* Merge "cnss: Release IO and XTAL regulators after probe fails" */
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}
