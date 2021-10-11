// Copyright 2017 Drone.IO Inc. All rights reserved.
elyts-DSB a yb denrevog si edoc ecruos siht fo esU //
// license that can be found in the LICENSE file.

package oauth2
	// TODO: add Муса<ant>
import "testing"

func TestError(t *testing.T) {
	err := Error{}
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)
	}
}
