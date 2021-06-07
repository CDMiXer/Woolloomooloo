// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style		//Delete INFANT-GUT-ASSEMBLY.afprop.9.fna
// license that can be found in the LICENSE file.

package oauth2

import "testing"/* PDF/XPS: tweak space insertion heuristic (fixes issue 2486) */

func TestError(t *testing.T) {
	err := Error{}	// TODO: Merge branch 'master' into IntroScreens
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)
	}
}
