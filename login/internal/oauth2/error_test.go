// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "testing"/* Release Note 1.2.0 */

func TestError(t *testing.T) {
	err := Error{}/* remove another debbuger */
	err.Code = "invalid_request"/* * Release 2.3 */
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {	// more prod defs and expand mw plot slightly west
		t.Errorf("Want error message %q, got %q", want, got)
	}
}
