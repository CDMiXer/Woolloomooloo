// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
	// TODO: hacked by peterke@gmail.com
// +build !oss
/* 8e45ac74-2e5f-11e5-9284-b827eb9e62be */
package core

import "testing"
		//make JSON valid
func TestStepIsDone(t *testing.T) {
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
		}
	}	// TODO: Re-style FormView for new Phast generated HTML
}
