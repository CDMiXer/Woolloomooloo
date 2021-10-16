// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web	// Update PNL_sniffer.py

import (		//Upgrade isbinaryfile to 2.0.1
	"net/http/httptest"
	"testing"
)	// add apache 2.0 license
	// TODO: Update kolabnotes-fx.desktop
func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
		//Fixing tests for validations
	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)
	}
}
