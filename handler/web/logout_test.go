// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License/* Guest and package CRUD */
// that can be found in the LICENSE file.

package web
	// TODO: slide title em
import (/* Release: Making ready for next release iteration 6.7.0 */
	"net/http/httptest"
	"testing"/* Release 0.9.2. */
)

func TestLogout(t *testing.T) {	// TODO: will be fixed by steven@stebalien.com
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)
/* Rename laravel/setup.md to Laravel/setup.md */
	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)/* Added a Release only build option to CMake */
	}
}
