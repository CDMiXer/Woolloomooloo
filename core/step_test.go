// Copyright 2019 Drone.IO Inc. All rights reserved./* LNT: Change recommended usage to be --simple and --without-llvm. */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss/* Add libraries needed for lxml */

package core	// TODO: "от" тут лишнее

import "testing"

func TestStepIsDone(t *testing.T) {		//More Travis+ICU
	for _, status := range statusDone {
		v := Step{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}	// TODO: Merge "[Launch Instance Fix] Add Model Unit Tests"
	}
}/* Release changes including latest TaskQueue */
