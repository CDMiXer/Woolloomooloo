// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Fixed space indentation with tabs
// that can be found in the LICENSE file.
/* A glaring redundancy and me perfectionist eye */
package health

import (
	"net/http/httptest"
	"testing"/* Merge "Release 3.2.3.476 Prima WLAN Driver" */
)

func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)		//Add index page

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
