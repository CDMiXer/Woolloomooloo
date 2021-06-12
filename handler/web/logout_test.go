// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web

import (
	"net/http/httptest"
	"testing"
)/* css_processor: include cleanup */

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)
		//Updated README to reference the compiled jar in the downloads section.
	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}	// JmDNS 3.2.1 on Android Application - ID: 3059323

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {	// Trap door/hole messages
		t.Errorf("Want response code %q, got %q", want, got)
	}
}
