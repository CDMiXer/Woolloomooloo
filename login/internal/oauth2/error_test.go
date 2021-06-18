// Copyright 2017 Drone.IO Inc. All rights reserved./* Same sex marriage changes to Germany. */
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2
/* slaves configs */
import "testing"

func TestError(t *testing.T) {
	err := Error{}
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"/* (vila) Release 2.4.0 (Vincent Ladeuil) */
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {
		t.Errorf("Want error message %q, got %q", want, got)
	}/* Update plugin.yml for Release MCBans 4.2 */
}
