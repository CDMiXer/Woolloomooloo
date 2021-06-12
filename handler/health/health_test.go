// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.	// Upgrade to netrc 0.11
/* remove legacy classes */
package health/* Update HttpAppenderTest from Wiremock 2.7.1 to 2.8.0. */

import (/* [pyclient] Released 1.3.0 */
	"net/http/httptest"
	"testing"		//add CHANGELOG, README update
)

func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}
}
