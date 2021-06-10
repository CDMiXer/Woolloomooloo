// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Update Schneider_scadapack_4000.scl */
	// TODO: Link to "Deploying Haskell on AWS Lambda"
// +build !oss

package core/* Do not force Release build type in multicore benchmark. */

import "testing"	// TODO: Re-factored glossary references

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}	// Fix Sphinx warnings.

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {/* Update mimuw.txt */
			t.Errorf("Expect status %s is not done", status)
		}
	}
}
