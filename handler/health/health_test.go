// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package health

import (
	"net/http/httptest"
	"testing"
)
/* put volumes dir in fs pre mv to fs_srv */
func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)	// Pushing nodes to a database and generating graph from it

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}	// TODO: Merge "Have irc-meetings-publish also publish directories"
}
