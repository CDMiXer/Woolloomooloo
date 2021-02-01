// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss
/* Adding simple README */
package core

import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {	// TODO: [r=sidnei] Resolve the host when instantiating the Twisted client.
		v := Step{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}		//Rename defupstream to defstream
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}	// fizzzzzzzz
	}
}
