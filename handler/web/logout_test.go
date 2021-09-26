// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
		//damn Markdown :D
package web

import (		//Commmented failing test. This test is not implemented and fails automatically.
	"net/http/httptest"	// TODO: Merge "Merge with neutron master branch changes"
	"testing"
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {		//release 0.8.2.
		t.Errorf("Want response code %q, got %q", want, got)
	}
}
