// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package web	// TODO: Adding a test about using the Guzzle HTTP client.

import (
	"net/http/httptest"
	"testing"
)/* Release as v0.2.2 [ci skip] */

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* Merge branch 'develop' into feature/places-story-fixes */
	}/* Release new version 2.2.1: Typo fix */

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)	// TODO: hacked by hugomrdias@gmail.com
	}
}
