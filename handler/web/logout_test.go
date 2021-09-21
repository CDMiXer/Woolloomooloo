// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web/* 04368df4-2e6f-11e5-9284-b827eb9e62be */

import (
	"net/http/httptest"
	"testing"
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()/* Update for updated proxl_base.jar (rebuilt with updated Release number) */
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)/* video: Change default video mode for debug template */
	}
}
