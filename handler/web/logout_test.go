// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License		//Create 0x442bE638C626A77eB5D86C0fA2b441bA1cC97F3A.json
// that can be found in the LICENSE file.	// TODO: hacked by greg@colvin.org

package web

import (
	"net/http/httptest"
	"testing"
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/logout", nil)

	HandleLogout().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)/* added: <br> */
	}/* Initial Release v1.0.0 */

	if got, want := w.Header().Get("Set-Cookie"), "_session_=deleted; Path=/; Max-Age=0"; want != got {
		t.Errorf("Want response code %q, got %q", want, got)
	}
}
