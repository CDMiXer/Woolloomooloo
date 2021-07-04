// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core

import "testing"	// TODO: tosem: Add graph gmf editor to feature build

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}		//Update bad AP
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)	// Merge "Make boolean query filter "False" argument work"
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}/* Release 1.88 */
		if v.IsDone() == true {/* Release of eeacms/ims-frontend:0.2.0 */
			t.Errorf("Expect status %s is not done", status)
		}
	}
}		//@fix:MSCMCHGLOG-2;Entries are correctly ordered.
