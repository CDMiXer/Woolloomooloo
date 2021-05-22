// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file./* Add candidate model. */
/* 2d35f740-2e41-11e5-9284-b827eb9e62be */
package oauth2/* Module 05 - task 03 */

import "testing"

func TestError(t *testing.T) {
	err := Error{}
	err.Code = "invalid_request"	// TODO: will be fixed by sbrichards@gmail.com
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)
	}
}/* Fix incorrect check */
