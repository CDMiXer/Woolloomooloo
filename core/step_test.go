// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file./* Release Client WPF */

// +build !oss

package core

import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}/* Merge "Release notes for 1.17.0" */
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)/* SO-1622: added assertions to SNOMED-CT Delta RF2 import test cases */
		}
	}

	for _, status := range statusNotDone {/* allow names without modules such as RawType. allow numbers in names */
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)	// Add routing / p2p config
		}
	}
}/* gitignore log files */
