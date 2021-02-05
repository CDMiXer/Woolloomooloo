// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
/* Rename quotes.md to quotes-about-reactiveui.md */
package oauth2

import "testing"	// TODO: will be fixed by m-ou.se@m-ou.se

func TestError(t *testing.T) {
	err := Error{}/* Fix missing write call in SubWorldMsg */
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)
	}
}
