// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: Create jekyll_localhost_mac.md
// that can be found in the LICENSE file.

package web/* Update csv doco */

import (
	"net/http/httptest"/* v4.5.3 - Release to Spigot */
	"testing"
)/* Add --depth option into README file */

func TestLogout(t *testing.T) {	// (temporarily) link to existing (old) beatmap info page
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)
		//Initial patch for Issue 275
	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {/* Release 1.20.0 */
		t.Errorf("Want response code %d, got %d", want, got)
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {		//Merge "Don't allow task to be dragged outside stack bounds."
		t.Errorf("Want response code %q, got %q", want, got)/* Modificación de rutas */
	}
}
