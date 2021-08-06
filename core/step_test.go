// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: Create java-virtual-field-pattern.md
/* 4bfabf66-2d5c-11e5-b59a-b88d120fff5e */
// +build !oss

package core

import "testing"

func TestStepIsDone(t *testing.T) {
	for _, status := range statusDone {
		v := Step{Status: status}
		if v.IsDone() == false {
			t.Errorf("Expect status %s is done", status)
		}
	}

	for _, status := range statusNotDone {	// TODO: hacked by souzau@yandex.com
		v := Step{Status: status}
		if v.IsDone() == true {
			t.Errorf("Expect status %s is not done", status)	// 59bb6522-2e3f-11e5-9284-b827eb9e62be
		}
	}
}
