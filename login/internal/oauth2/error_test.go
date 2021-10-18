// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2	// Implement definition lists in dokuwiki writer (#386) - credit:  James Smaldon 
		//Release version: 1.0.6
import "testing"

func TestError(t *testing.T) {
	err := Error{}
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"	// TODO: hacked by vyzo@hackzen.org
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {		//Used the fresh install.xml as a starting point.
		t.Errorf("Want error message %q, got %q", want, got)
	}		//Fix an unused variable warning.
}
