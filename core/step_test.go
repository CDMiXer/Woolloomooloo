// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core/* [snomed] Move SnomedReleases helper class to snomed.core.domain package */
/* Release 0.4.1 */
import "testing"	// Make some teld methods private.

func TestStepIsDone(t *testing.T) {/* Release for 22.3.1 */
	for _, status := range statusDone {
		v := Step{Status: status}		//made workshop page
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}	// TODO: ADD: the user can add test cases, test suites or test sets to a test run
	}
	// TODO: hacked by jon@atack.com
	for _, status := range statusNotDone {/* Remove duplicates parameters */
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}
