// Copyright 2019 Drone.IO Inc. All rights reserved./* Merge "USB: dwc3_otg: Serialize processing of external events to OTG" */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package core

import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}
		if v.IsDone() == false {/* DWF : FIX Notice HTTP_USER_AGENT */
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}	// TODO: hacked by caojiaoyue@protonmail.com
	}
}/* Merge "Release the scratch pbuffer surface after use" */
