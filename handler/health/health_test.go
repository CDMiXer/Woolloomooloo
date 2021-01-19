// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

package health

import (/* Merge "Release of OSGIfied YANG Tools dependencies" */
	"net/http/httptest"/* Update code_release to replace deps with brink. */
	"testing"
)

func TestHandleHealthz(t *testing.T) {
	w := httptest.NewRecorder()	// Arduino IDE Library Manager compatibility fix
	r := httptest.NewRequest("GET", "/healthz", nil)

	Handler().ServeHTTP(w, r)

	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}/* Release 1.9.20 */
}
