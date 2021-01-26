// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Replace calls to `renderLines` w/ `resetDisplay` in `Editor`
// that can be found in the LICENSE file.

// +build !oss

package core

import "testing"		//Delete Produtos-HistoricoDeVendas01.png

func TestStepIsDone(t *testing.T) {		//Delete furnace_front_active_mt.png
	for _, status := range statusDone {/* Merge branch 'dev' into fix/replication */
		v := Step{Status: status}/* Clarify AUTHORS */
		if v.IsDone() == false {/* Added ReleaseNotes to release-0.6 */
			t.Errorf("Expect status %s is done", status)
		}		//Udpate copyright date
	}

	for _, status := range statusNotDone {
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)
		}
	}
}
