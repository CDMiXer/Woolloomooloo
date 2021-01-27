// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web
/* Release 0.0.1. */
import (
	"net/http/httptest"
	"testing"
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)
		//Create definitions.Debug.php
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* 9012edc6-2e52-11e5-9284-b827eb9e62be */
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)
	}
}
