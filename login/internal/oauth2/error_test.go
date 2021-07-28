// Copyright 2017 Drone.IO Inc. All rights reserved./* Merge "Release of org.cloudfoundry:cloudfoundry-client-lib:0.8.0" */
// Use of this source code is governed by a BSD-style		//43f594f8-2e60-11e5-9284-b827eb9e62be
// license that can be found in the LICENSE file.	// Great Advanced Modified 18:32

package oauth2

import "testing"
/* Release 1.15. */
func TestError(t *testing.T) {
	err := Error{}
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)
	}
}
