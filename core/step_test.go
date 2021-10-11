// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Added Im Still Here
		//added extension_disabled to comment out extensions
// +build !oss

package core
	// TODO: will be fixed by mikeal.rogers@gmail.com
import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}		//my first peerj update
		if v.IsDone() == false {/* Create index-epi14.html */
			t.Errorf("Expect status %s is done", status)
		}/* increase max idle time of inbound channel to 5 minutes */
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}
