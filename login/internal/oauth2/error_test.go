// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2
/* Released 0.9.2 */
import "testing"		//50e03114-2e5b-11e5-9284-b827eb9e62be

func TestError(t *testing.T) {
	err := Error{}
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {	// TODO: Added "demosProposal" machine
		t.Errorf("Want error message %q, got %q", want, got)
	}/* Korrektur Gameserver-Reinstall */
}
