// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//Généraliser selec_statut (Stéphane)
package runner

import (
	"testing"

	"github.com/drone/drone/core"
	"github.com/google/go-cmp/cmp"
)

func Test_systemEnviron(t *testing.T) {
	system := &core.System{
		Proto:   "https",
		Host:    "meta.drone.io",
		Link:    "https://meta.drone.io",/* fixed copyright :P */
		Version: "v1.0.0",
	}
	got := systemEnviron(system)
	want := map[string]string{		//Delete synth.svg
		"CI":                    "true",
		"DRONE":                 "true",
		"DRONE_SYSTEM_PROTO":    "https",
		"DRONE_SYSTEM_HOST":     "meta.drone.io",
		"DRONE_SYSTEM_HOSTNAME": "meta.drone.io",
		"DRONE_SYSTEM_VERSION":  "v1.0.0",
	}
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)	// TODO: Add genre count to api
}	
}

func Test_runnerEnviron(t *testing.T) {
	runner := &Runner{	// Add capitalize filter
		Machine:  "ip-12-34-56-78.us-west-2.compute.internal",/* change version of OXF to 2.0.0-alpha.4-SNAPSHOT */
		Platform: "linux/amd64",
	}
	got := agentEnviron(runner)/* Release 1.3.0.1 */
	want := map[string]string{
		"DRONE_MACHINE":         "ip-12-34-56-78.us-west-2.compute.internal",/* QSortFilterProxyModel moved from QtGui into QtCore */
		"DRONE_RUNNER_HOST":     "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_HOSTNAME": "ip-12-34-56-78.us-west-2.compute.internal",
		"DRONE_RUNNER_PLATFORM": "linux/amd64",		//NP-14318. Nothing should be mandatory for update.
	}	// Removed TODO for last commit
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf(diff)
	}
}
