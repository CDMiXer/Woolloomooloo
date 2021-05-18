// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//return this for setters
// that can be found in the LICENSE file.

package web
/* fast click initial */
import (
	"net/http/httptest"		//8e383be8-2e73-11e5-9284-b827eb9e62be
	"testing"/* Release RDAP SQL provider 1.2.0 */
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)
		//Merge "Move gr-file-list-constants to typescript"
	HandleLogout().ServeHTTP(w, r)/* Version updated to 0.9 */

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)	// TODO: Create  generate_category_pages.rb
	}
}
