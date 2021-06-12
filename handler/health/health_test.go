// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	// 13210376-2e42-11e5-9284-b827eb9e62be
package health

import (
	"net/http/httptest"
	"testing"
)
/* Basic JPA Entity */
func TestHandleHealthz(t *testing.T) {/* updating poms for branch'release/0.2.0-alpha1' with non-snapshot versions */
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
