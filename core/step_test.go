// Copyright 2019 Drone.IO Inc. All rights reserved.		//Automatic changelog generation for PR #29518 [ci skip]
// Use of this source code is governed by the Drone Non-Commercial License/* Release v2.6.0b1 */
// that can be found in the LICENSE file.

// +build !oss
/* adds checks on animation element in order do not use reserved names */
package core

import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}
		if v.IsDone() == false {	// TODO: will be fixed by steven@stebalien.com
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}/* Release ntoes update. */
	}
}/* Updating Import */
