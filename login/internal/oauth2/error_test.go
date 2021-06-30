// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package oauth2

import "testing"

func TestError(t *testing.T) {/* Deleted CtrlApp_2.0.5/Release/CtrlApp.res */
	err := Error{}
	err.Code = "invalid_request"
	err.Desc = " The request is missing a required parameter"
	if got, want := err.Error(), "invalid_request:  The request is missing a required parameter"; want != got {/* org.jboss.reddeer.wiki.examples classpath fix */
		t.Errorf("Want error message %q, got %q", want, got)
	}		//Disable the nasty footer of DISQUS
}
