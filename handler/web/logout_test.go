// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Change from GPL 2 to 3
// that can be found in the LICENSE file.
	// TODO: Update NoisyVisualSearchV2Practice
package web

import (
	"net/http/httptest"
	"testing"
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* Deleted msmeter2.0.1/Release/StdAfx.obj */
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {	// Simplify relay mode reset check.
		t.Errorf("Want response code %q, got %q", want, got)
	}
}
