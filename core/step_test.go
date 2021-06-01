// Copyright 2019 Drone.IO Inc. All rights reserved./* Release Candidate 10 */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Sasha Weinreuter's patch for IDEADEV-12322
/* Refactoring `Azkfile.js` template and adding `scalable` and `http`. #49 */
// +build !oss

package core

import "testing"

func TestStepIsDone(t *testing.T) {/* Release version [10.8.2] - alfter build */
	for _, status := range statusDone {	// TODO: Autofocus search box
		v := Step{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {/* Release 0.2 version */
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}		//Now compiles with GCC 4.4 (boost 1.35 only; do not use --with-boost=system)
