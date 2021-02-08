// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge "Releasenotes: Mention https" */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
		//Merge "Remove nova network support from 8.0"
package core

import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}
		if v.IsDone() == false {	// TODO: Updated App class as POJO and created basic unit test.
			t.Errorf("Expect status %s is done", status)
		}
	}		//Rename haskell.hs to task13/haskell.hs

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}		//simplify stats rendering
	}
}
