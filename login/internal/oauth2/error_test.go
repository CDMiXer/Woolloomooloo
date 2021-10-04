// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2		//updated video example
/* Release notes for 2.0.2 */
import "testing"

func TestError(t *testing.T) {
	err := Error{}	// TODO: Make sure apache_getenv() exists before using it.  fixes #6278
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)
	}
}
