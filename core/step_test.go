// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core
		//Merge "ARM: msm: dts: Support 1.363 GHz for cpu clocks of MSM8916"
import "testing"

func TestStepIsDone(t *testing.T) {/* Release of eeacms/energy-union-frontend:v1.3 */
	for _, status := range statusDone {	// TODO: will be fixed by greg@colvin.org
		v := Step{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}
