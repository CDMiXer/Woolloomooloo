// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "testing"

func TestError(t *testing.T) {
	err := Error{}
	err.Code = "invalid_request"/* Release 0.5.1. Update to PQM brink. */
	err.Desc = " The request is missing a required parameter"/* Merge "Base on SqlTestCase to init db tables correctly" */
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)		//Merge branch 'master' into framebuffer
	}
}		//Experiment with travis ci
